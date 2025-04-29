package sessionstorage

import (
	"codeproc/infrastructure/repository"
	"time"
)

type sessionstorage struct {
	session map[string]*session
}

type session struct {
	uid          string
	timeAccessed time.Time
}

type userstorage struct {
	user map[string]*data
}

type data struct {
	login    string
	password string
}

func NewSessionStorage() *sessionstorage {
	return &sessionstorage{make(map[string]*session)}
}

func NewObject() *userstorage {
	return &userstorage{make(map[string]*data)}
}

func (s *sessionstorage) Post(sid, uid string) {
	s.session[sid] = &session{
		uid:          uid,
		timeAccessed: time.Now(),
	}
}

func (s *sessionstorage) Get(sid string) error {
	if _, ok := s.session[sid]; !ok {
		return repository.Unauthorized
	}
	return nil
}

func (s *sessionstorage) Delete(sid string) {
	delete(s.session, sid)
}

func (s *sessionstorage) SessionGC(maxlifetime int64) {
	for sid, _ := range s.session {
		if s.session[sid].timeAccessed.Unix()+maxlifetime < time.Now().Unix() {
			s.Delete(sid)
		}
	}
}

func (u *userstorage) CheckUser(login, password string) (string, bool, bool) {
	var okLogin, okPassword bool = false, false
	for id, user := range u.user {
		if user.login == login {
			okLogin = true
		}
		if user.password == password {
			okPassword = true
		}
		if user.login == login && user.password == password {
			return id, okLogin, okPassword
		}
	}
	return "", okLogin, okPassword
}

func (u *userstorage) SignUp(login, password, uid string) error {
	if _, okLogin, _ := u.CheckUser(login, password); okLogin {
		return repository.LoginExist
	}
	u.user[uid] = &data{
		login:    login,
		password: password,
	}
	return nil
}

func (u *userstorage) SignIn(login, password string) (string, error) {
	if uid, okLogin, okPassword := u.CheckUser(login, password); okLogin && okPassword {
		return uid, nil
	}
	return "", repository.NotFound
}
