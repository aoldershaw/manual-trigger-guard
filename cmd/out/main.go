package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	. "github.com/aoldershaw/manual-trigger-guard"
	"github.com/google/go-github/v38/github"
	"golang.org/x/oauth2"
)

func main() {
	var input struct {
		Source Source `json:"source"`
		Params Source `json:"params"`
	}
	if err := json.NewDecoder(os.Stdin).Decode(&input); err != nil {
		log.Fatal(err)
	}
	source := input.Source.MergeWith(input.Params)
	err := validate(source, os.Getenv("BUILD_CREATED_BY"))
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(os.Stdout).Encode(map[string]interface{}{
		"version": Version{"v"},
	})
}

func validate(source Source, username string) error {
	if username == "" {
		return errors.New("Build was not manually triggered, or username is not available (did you set `expose_build_created_by: true`?)")
	}

	for _, allowedUser := range source.Users {
		if username == allowedUser {
			return nil
		}
	}

	ctx := context.Background()
	client := githubClient(ctx, source)

	for _, allowedTeam := range source.Teams {
		parts := strings.Split(allowedTeam, "/")
		if len(parts) != 2 {
			return fmt.Errorf("Improperly formatted team %q (expecting ORG/TEAM)", allowedTeam)
		}

		membership, _, err := client.Teams.GetTeamMembershipBySlug(ctx, parts[0], parts[1], username)
		if err != nil {
			log.Printf("Failed to check membership to team %v: %v", parts, err)
			continue
		}

		if membership.GetState() != "active" {
			continue
		}

		return nil
	}

	return fmt.Errorf("User %q is not in the list of allowed users/teams", username)
}

func githubClient(ctx context.Context, source Source) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: source.AccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}
