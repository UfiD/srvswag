package types

import (
	"codeproc/internal/domain"
	"codeproc/internal/infrastructure/repository"
	"encoding/json"
	"fmt"
	"net/http"
)

type PostObjectHandlerRequest struct {
	domain.Object
}

func CreatePostObjectHandlerRequest(r *http.Request) (*PostObjectHandlerRequest, error) {
	var req PostObjectHandlerRequest
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&req); err != nil {
		return nil, fmt.Errorf("error while decode json: %v", err)
	}
	return &req, nil
}

type GetObjectHandlerRequest struct {
	domain.Task
}

func CreateGetObjectHandlerRequest(r *http.Request) (*GetObjectHandlerRequest, error) {
	var req GetObjectHandlerRequest
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&req); err != nil {
		return nil, fmt.Errorf("error while decoding json: %v", err)
	}
	return &req, nil
}

type GetObjectHandlerResponse struct {
	ID     string `json:"task_id,omitempty"`
	Status string `json:"status,omitempty"`
	Result string `json:"result,omitempty"`
	Sid    string `json:"Authorization,omitempty"`
}

type PostAuthObjectHandlerRequest struct {
	domain.Userdata
}

func CreatePostAuthObjectHandlerRequest(r *http.Request) (*PostAuthObjectHandlerRequest, error) {
	var req PostAuthObjectHandlerRequest
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&req); err != nil {
		return nil, fmt.Errorf("error while decoding json: %v", err)
	}
	return &req, nil
}

func CreateAuthHeader(r *http.Request) (string, error) {
	req := r.Header.Get("Authorization")
	if req == "" {
		return req, fmt.Errorf("not found session")
	}
	return req, nil
}

func ProcessError(w http.ResponseWriter, err error, resp any) {
	if err == repository.NotFound {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	if err == repository.LoginExist {
		http.Error(w, "User with this username already exists", http.StatusBadRequest)
	} else if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
	}
}
