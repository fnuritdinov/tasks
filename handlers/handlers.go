package handlers

import (
	"TaskManager/internal/models"
	"TaskManager/internal/service"
	errs "TaskManager/pkg/errors"
	"TaskManager/pkg/logger"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"go.uber.org/zap"
)

type TaskHandler struct {
	service service.TaskService
	logger  logger.Logger
}

func NewTaskHandlers(service service.TaskService, logger logger.Logger) TaskHandler {
	return TaskHandler{
		service: service,
		logger:  logger,
	}
}

type taskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	IsDeleted   bool   `json:"isDeleted"`
}

func (t *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req taskRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		t.logger.Error("error from json.NewDecoder",
			zap.Error(err))

		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	err = t.service.Create(r.Context(), models.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	})
	if err != nil {
		if errors.Is(err, errs.ErrFromValidate) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		t.logger.Error("error from t.service.Create",
			zap.Error(err))
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	_ = json.NewEncoder(w).Encode(map[string]string{
		"message": "task created successfully",
	})

}

func (t *TaskHandler) Get(w http.ResponseWriter, r *http.Request) {

	tasks, err := t.service.Get(r.Context())
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		t.logger.Error("error from t.service.Get")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(tasks)
}

func (t *TaskHandler) GetWithFilter(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	tasks, err := t.service.GetWithFilters(r.Context(), models.TaskFilter{
		Status: status,
		Page:   page,
		Limit:  limit,
	})
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		t.logger.Error("error from t.service.Get")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(tasks)
}

func (t *TaskHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	task, err := t.service.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, errs.ErrFromValidate) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if errors.Is(err, errs.ErrNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		t.logger.Error("error from t.service.GetByID")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(task)

}

func (t *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var req taskRequest

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = t.service.Update(r.Context(), id, models.Tasks{
		Description: req.Description,
		Title:       req.Title,
		Status:      req.Status,
		IsDeleted:   req.IsDeleted,
	})
	if err != nil {
		if errors.Is(err, errs.ErrFromValidate) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if errors.Is(err, errs.ErrNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		t.logger.Error("error from t.service.Update")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(map[string]string{
		"message": "task updated successfully",
	})
}

func (t *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if id < 1 {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = t.service.Delete(r.Context(), id)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		t.logger.Error("error from t.service.Delete")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}
