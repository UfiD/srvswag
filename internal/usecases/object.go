package usecases

import "codeproc/internal/domain"

type Object interface {
	Post(object domain.Object) string
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
