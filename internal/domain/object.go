package domain

type Object struct {
	ID       string `json:"id,omitempty"`
	Code     string `json:"task" example:"your code"`
	Compiler string `json:"compiler" example:"python3"`
}

type ObjectResult struct {
	ID       string `json:"id"`
	Code     string `json:"task"`
	Compiler string `json:"compiler"`
	Result   string `json:"result"`
	Error    string `json:"error"`
}

type Userdata struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Task struct {
	ID string `json:"task_id"`
}
