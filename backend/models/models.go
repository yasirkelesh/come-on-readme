package models

type GeneralContent struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	GitURL string `json:"git_url"`
}

type File struct {
	Content  string `json:"content"`
	Encoding string `json:"encoding"`
}
