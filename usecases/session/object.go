package session

import (
	"codeproc/infrastructure/repository"
	"codeproc/pkg/uuid"
	"sync"
	"time"
)

type Manager struct {
	store       repository.Userstorage
	provide     repository.Provider
	lock        sync.Mutex
	maxlifetime int64
}

func NewObject(store repository.Userstorage, provide repository.Provider, maxlifetime int64) *Manager {
	return &Manager{
		store:       store,
		provide:     provide,
		maxlifetime: maxlifetime,
	}
}

func (m *Manager) SignUp(login, password string) error {
	uid := uuid.GetUUID()
	if err := m.store.SignUp(login, password, uid); err != nil {
		return err
	}
	return nil
}

func (m *Manager) SessionStart(login, password string) (string, error) {
	uid, err := m.store.SignIn(login, password)
	if err != nil {
		return "", err
	}
	sid := uuid.GetUUID()
	go m.provide.Post(sid, uid)
	return sid, nil
}

func (m *Manager) SessionRead(sid string) error {
	if err := m.provide.Get(sid); err != nil {
		return err
	}
	return nil
}

func (m *Manager) SessionDestroy(sid string) {
	m.provide.Delete(sid)
}

func (m *Manager) GC() {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.provide.SessionGC(m.maxlifetime)
	time.AfterFunc(time.Duration(m.maxlifetime), func() { m.GC() })
}
