package models

type GeneralContent struct {
	Name string `json:"name"`
	Type string `json:"type"`
	URL  string `json:"url"`
}

type FileContent struct {
	Content  string `json:"content"`
	Encoding string `json:"encoding"`
}

