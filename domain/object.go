package domain

type Object struct {
	Code     string `json:"task" example:"your code"`
	Compiler string `json:"compiler" example:"python3"`
}
