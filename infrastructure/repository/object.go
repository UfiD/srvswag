package repository

type Object interface {
	Post(id, code, compiler string)
	GetStatus(id string) (string, error)
	GetResult(id string) (string, error)
	Commit(id, result string)
}
