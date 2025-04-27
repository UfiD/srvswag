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
}

// New creates a new instance of Controller
func New(uc usecases.Object) *Controller {
	return &Controller{
		uc: uc,
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
	req, err := types.CreatePostObjectHandlerRequest(r)
	if err != nil {
		http.Error(w, "Bad request in post", http.StatusBadRequest)
		return
	}
	id := c.uc.Post(req.Code, req.Compiler)
	types.ProcessError(w, err, &types.GetObjectHandlerResponse{ID: id, Status: "", Result: ""})
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
	req, err := types.CreateGetObjectHandlerRequest(r)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	status, err := c.uc.GetStatus(req.ID)
	types.ProcessError(w, err, &types.GetObjectHandlerResponse{ID: "", Status: status, Result: ""})
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
	req, err := types.CreateGetObjectHandlerRequest(r)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	result, err := c.uc.GetResult(req.ID)
	types.ProcessError(w, err, &types.GetObjectHandlerResponse{ID: "", Status: "", Result: result})
}

// WithObjectHandlers registers object-related HTTP handlers.
func (c *Controller) WithObjectHandler(r chi.Router) {
	r.Route("/task", func(r chi.Router) {
		r.Post("/", c.post)
		r.Get("/status", c.getStatus)
		r.Get("/result", c.getResult)
	})
}
