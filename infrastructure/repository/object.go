package repository

type Object interface {
	Post(id, code, compiler string)
	GetStatus(id string) (string, error)
	GetResult(id string) (string, error)
	Put(id, result string)
}

type Provider interface {
	Post(sid, uid string)
	Get(sid string) error
	Delete(sid string)
	SessionGC(maxlifetime int64)
}

type Userstorage interface {
	SignUp(login, password, uid string) error
	SignIn(login, password string) (string, error)
}
