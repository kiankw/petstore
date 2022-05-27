package model

type Pet struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name"`
}
