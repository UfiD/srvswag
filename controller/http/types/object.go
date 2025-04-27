package types

import (
	"codeproc/domain"
	"codeproc/infrastructure/repository"
	"encoding/json"
	"fmt"
	"net/http"
)

type PostObjectHandlerRequest struct {
	domain.Object
}

func CreatePostObjectHandlerRequest(r *http.Request) (*PostObjectHandlerRequest, error) {
	var req PostObjectHandlerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("error while decode json: %v", err)
	}
	return &req, nil
}

type GetObjectHandlerRequest struct {
	domain.Task
}

func CreateGetObjectHandlerRequest(r *http.Request) (*GetObjectHandlerRequest, error) {
	var req GetObjectHandlerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("error while decode json: %v", err)
	}
	return &req, nil
}

type GetObjectHandlerResponse struct {
	ID     string `json:"task_id,omitempty"`
	Status string `json:"status,omitempty"`
	Result string `json:"result,omitempty"`
}

func ProcessError(w http.ResponseWriter, err error, resp any) {
	if err == repository.NotFound {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {

	}
}
