package guard

type Source struct {
	AccessToken string   `json:"access_token"`
	Teams       []string `json:"teams"`
	Users       []string `json:"users"`
}

func (s Source) MergeWith(s2 Source) Source {
	out := s
	if s2.AccessToken != "" {
		out.AccessToken = s2.AccessToken
	}
	if s2.Teams != nil {
		out.Teams = s2.Teams
	}
	if s2.Users != nil {
		out.Users = s2.Users
	}
	return out
}

type Version struct {
	Version string `json:"version"`
}
