package handler

type Task struct {
	Title  string `json:"title"`
	Done   bool   `json:"done"`
	DoneAt string `json:"done_at,omitempty"`
}

