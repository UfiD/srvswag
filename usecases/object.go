package usecases

type Object interface {
	Post(code, compiler string) string
	GetStatus(id string) (string, error)
	GetResult(id string) (string, error)
}

type Manager interface {
	SignUp(login, password string) error
	SessionStart(login, password string) (string, error)
	SessionRead(sid string) error
	SessionDestroy(sid string)
	GC()
}
