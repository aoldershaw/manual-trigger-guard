package main

import (
	"encoding/json"
	"os"

	. "github.com/aoldershaw/manual-trigger-guard"
)

func main() {
	json.NewEncoder(os.Stdout).Encode([]Version{
		{"v"},
	})
}
