package handlers

import (
	"TaskManager/internal/models"
	"TaskManager/internal/service"
	"TaskManager/pkg/logger"
	"encoding/json"
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
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	log := t.logger.With(zap.String("handler", "Create"))
	err = t.service.Create(r.Context(), models.Tasks{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	})
	if err != nil {
		handleError(w, log, err)
		return
	}

	w.WriteHeader(http.StatusCreated)

	_ = json.NewEncoder(w).Encode(map[string]string{
		"message": "task created successfully",
	})

}

func (t *TaskHandler) Get(w http.ResponseWriter, r *http.Request) {

	log := t.logger.With(zap.String("handler", "Get"))
	tasks, err := t.service.Get(r.Context())
	if err != nil {
		handleError(w, log, err)
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

	log := t.logger.With(zap.String("handler", "GetWithFilter"))

	tasks, err := t.service.GetWithFilters(r.Context(), models.TaskFilter{
		Status: status,
		Page:   page,
		Limit:  limit,
	})
	if err != nil {
		handleError(w, log, err)
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

	log := t.logger.With(zap.String("handler", "GetByID"))

	task, err := t.service.GetByID(r.Context(), id)
	if err != nil {
		handleError(w, log, err)
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

	log := t.logger.With(zap.String("handler", "Update"))

	err = t.service.Update(r.Context(), id, models.Tasks{
		Description: req.Description,
		Title:       req.Title,
		Status:      req.Status,
		IsDeleted:   req.IsDeleted,
	})
	if err != nil {
		handleError(w, log, err)
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

	log := t.logger.With(zap.String("handler", "Delete"))

	err = t.service.Delete(r.Context(), id)
	if err != nil {
		handleError(w, log, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}
