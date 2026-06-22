package handlers

import "net/http"

func New(handler *TaskHandler) *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("POST /tasks", handler.Create)
	mux.HandleFunc("GET /tasks", handler.Get)
	mux.HandleFunc("GET /tasks/{id}", handler.GetByID)
	mux.HandleFunc("PUT /tasks/{id}", handler.Update)
	mux.HandleFunc("DELETE /tasks/{id}", handler.Delete)
	mux.HandleFunc("GET /tasks?status=done?page=1&limit=10", handler.GetWithFilter)
	mux.HandleFunc("PATCH /tasks/{id}/status", handler.Completed)
	mux.HandleFunc("POST /users", handler.CreateUser)
	mux.HandleFunc("GET /users", handler.GetUsers)
	mux.HandleFunc("GET /users/{id}", handler.GetUserByID)
	mux.HandleFunc("PUT /users/{id}", handler.UpdateUser)
	mux.HandleFunc("DELETE /users/{id}", handler.DeleteUser)

	return mux
}
