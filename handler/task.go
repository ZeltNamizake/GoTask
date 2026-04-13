package handler

type Task struct {
	Title    string `json:"title"`
	Done     bool   `json:"done"`
	CreateAt string `json:"create_at,omitempty"`
	DoneAt   string `json:"done_at,omitempty"`
}
