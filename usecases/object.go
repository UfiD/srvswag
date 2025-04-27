package usecases

type Object interface {
	//Register(login, password string)
	//Auth(login, password string)
	Post(code, compiler string) string
	GetStatus(id string) (string, error)
	GetResult(id string) (string, error)
}
