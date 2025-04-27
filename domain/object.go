package domain

type Object struct {
	Code     string `json:"task"`
	Compiler string `json:"compiler"`
}

type Task struct {
	ID string `json:"task_id"`
}
