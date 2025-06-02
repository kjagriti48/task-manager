package models

type Task struct {
	ID        string `json:"id,omitempty"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	User      string `json:"user,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
}
