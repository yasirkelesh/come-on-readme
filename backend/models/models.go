package models

type Content []struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	GitURL string `json:"git_url"`
}
