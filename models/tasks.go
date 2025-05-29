package models

type Task struct {
	ID        string `json:"id`
	Title     string `json:"title"`
	Completed bool   `json:"completed`
	DueDate   string `json:"dueDate"`
	User      string `json:"user"`
	CreatedAt string `json:"createdAt"` //
}
