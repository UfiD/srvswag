package ram_storage

import "codeproc/infrastructure/repository"

type task struct {
	code     string
	compiler string
	status   string
	result   string
}

type storage struct {
	task map[string]*task
}

func NewStorage() *storage {
	return &storage{
		make(map[string]*task),
	}
}

func newTask(code, compiler string) *task {
	return &task{
		code:     code,
		compiler: compiler,
		status:   "In progress",
		result:   "Please wait",
	}
}

func (s *storage) Post(id, code, compiler string) {
	task := newTask(code, compiler)
	s.task[id] = task
}

func (s *storage) GetStatus(id string) (string, error) {
	_, ok := s.task[id]
	if !ok {
		return "", repository.NotFound
	}
	return s.task[id].status, nil
}

func (s *storage) GetResult(id string) (string, error) {
	_, ok := s.task[id]
	if !ok {
		return "", repository.NotFound
	}
	return s.task[id].result, nil
}

func (s *storage) Put(id, result string) {
	s.task[id].status = "Ready"
	s.task[id].result = result
}
