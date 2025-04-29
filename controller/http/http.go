package controller

import (
	"codeproc/controller/http/types"
	"codeproc/usecases"
	"net/http"

	"github.com/go-chi/chi"
)

// Object represents an HTTP handler for managing objects.
type Controller struct {
	uc usecases.Object
	m  usecases.Manager
}

// New creates a new instance of Controller
func New(uc usecases.Object, m usecases.Manager) *Controller {
	return &Controller{
		uc: uc,
		m:  m,
	}
}

// @Summary Отправка кода и названия языка программирования
// @Description Создание новой задачи, старт работы обработчика.
// @Description Возвращает ID задачи
// @Tags object
// @Accept  json
// @Produce json
// @Param Object body types.PostObjectHandlerRequest true "Task and language name"
// @Success 200 {string} string "Код успешно загружен"
// @Failure 400 {string} string "Bad request"
// @Router / [post]
func (c *Controller) post(w http.ResponseWriter, r *http.Request) {
	sid, err := types.CreateAuthHeader(r)
	if err != nil {
		http.Error(w, "Session token is invalid", http.StatusUnauthorized)
		return
	}
	err = c.m.SessionRead(sid)
	if err != nil {
		http.Error(w, "Session token is invalid", http.StatusUnauthorized)
		return
	}
	req, err := types.CreatePostObjectHandlerRequest(r)
	if err != nil {
		http.Error(w, "Bad request in post", http.StatusBadRequest)
		return
	}
	id := c.uc.Post(req.Code, req.Compiler)
	types.ProcessError(w, err, &types.GetObjectHandlerResponse{ID: id})
}

// @Summary Проверка статуса выполнения
// @Description Проверяет статус выполнения задачи с указанным ID
// @Description Возвращает статус выполнения задачи с указанным ID
// @Tags object
// @Accept  json
// @Produce json
// @Param ID query string true "task_id"
// @Success 200 {string} types.GetObjectHandlerResponse
// @Failuer 404 {string} string "Object not found"
// @Failure 400 {string} string "Bad request"
// @Router /status [get]
func (c *Controller) getStatus(w http.ResponseWriter, r *http.Request) {
	sid, err := types.CreateAuthHeader(r)
	if err != nil {
		http.Error(w, "Session token is invalid", http.StatusUnauthorized)
		return
	}
	err = c.m.SessionRead(sid)
	if err != nil {
		http.Error(w, "Session token is invalid", http.StatusUnauthorized)
		return
	}
	req, err := types.CreateGetObjectHandlerRequest(r)
	if err != nil {
		http.Error(w, "Bad request u are pussy", http.StatusBadRequest)
		return
	}
	status, err := c.uc.GetStatus(req.ID)
	types.ProcessError(w, err, &types.GetObjectHandlerResponse{Status: status})
}

// @Summary Получение результата
// @Description Возвращает результат задачи с указанным ID
// @Tags object
// @Accept  json
// @Produce json
// @Param ID query string true "task_id"
// @Success 200 {string} types.GetObjectHandlerResponse
// @Failuer 404 {string} string "Object not found"
// @Failure 400 {string} string "Bad request"
// @Router /result [get]
func (c *Controller) getResult(w http.ResponseWriter, r *http.Request) {
	sid, err := types.CreateAuthHeader(r)
	if err != nil {
		http.Error(w, "Session token is invalid", http.StatusUnauthorized)
		return
	}
	err = c.m.SessionRead(sid)
	if err != nil {
		http.Error(w, "Session token is invalid", http.StatusUnauthorized)
		return
	}
	req, err := types.CreateGetObjectHandlerRequest(r)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	result, err := c.uc.GetResult(req.ID)
	types.ProcessError(w, err, &types.GetObjectHandlerResponse{Result: result})
}

func (c *Controller) signUp(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreatePostAuthObjectHandlerRequest(r)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	err = c.m.SignUp(req.Login, req.Password)
	types.ProcessError(w, err, nil)
}

func (c *Controller) signIn(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreatePostAuthObjectHandlerRequest(r)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	sid, err := c.m.SessionStart(req.Login, req.Password)
	types.ProcessError(w, err, &types.GetObjectHandlerResponse{Sid: sid})
}

// WithObjectHandlers registers object-related HTTP handlers.
func (c *Controller) WithObjectHandler(r chi.Router) {

	r.Post("/signup", c.signUp)
	r.Post("/signin", c.signIn)

	r.Route("/task", func(r chi.Router) {
		r.Post("/", c.post)
		r.Get("/status", c.getStatus)
		r.Get("/result", c.getResult)
	})
}
