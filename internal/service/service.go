package service

import (
	"TaskManager/internal/models"
	"TaskManager/internal/repository"
	"context"
)

type TaskService interface {
	Create(ctx context.Context, task models.Tasks) error
	Get(ctx context.Context) ([]models.Tasks, error)
	GetWithFilters(ctx context.Context, filter models.TaskFilter) ([]models.Tasks, error)
	GetByID(ctx context.Context, id int) (models.Tasks, error)
	Update(ctx context.Context, id int, updatedTasks models.Tasks) error
	Compiled(ctx context.Context, id int, task models.Tasks) (models.Tasks, error)

	Delete(ctx context.Context, id int) error
	CreateUser(ctx context.Context, user models.UserReq) error
	GetUsers(ctx context.Context) ([]models.UserReq, error)
	GetUserByID(ctx context.Context, id int) (models.UserReq, error)
	UpdateUser(ctx context.Context, id int, updatedUser models.UserReq) error
	DeleteUser(ctx context.Context, id int) error
}

type Service struct {
	Repo repository.TaskRepo
}

func New(Repo repository.TaskRepo) TaskService {
	return &Service{
		Repo: Repo,
	}
}
