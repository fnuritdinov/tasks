package repository

import (
	"TaskManager/internal/models"
	"context"

	"gorm.io/gorm"
)

type TaskRepo interface {
	Create(ctx context.Context, task models.Tasks) error
	Get(ctx context.Context) ([]models.Tasks, error)
	GetWithFilters(ctx context.Context, filter models.TaskFilter) ([]models.Tasks, error)
	GetByID(ctx context.Context, id int) (models.Tasks, error)
	Update(ctx context.Context, id int, updatedData models.Tasks) error
	Delete(ctx context.Context, id int) error
	Compiled(ctx context.Context, id int, task models.Tasks) (models.Tasks, error)
	SelectTask(ctx context.Context, id int) (models.Tasks, error)

	CreateUser(ctx context.Context, user models.UserReq) error
	GetUsers(ctx context.Context) ([]models.UserReq, error)
	GetUserByID(ctx context.Context, id int) (models.UserReq, error)
	UpdateUser(ctx context.Context, id int, updatedUser models.UserReq) error
	DeleteUser(ctx context.Context, id int) error
	SelectUser(ctx context.Context, userID int) (models.User, error)
}

type Repo struct {
	DB *gorm.DB
}

func New(DB *gorm.DB) TaskRepo {
	return &Repo{
		DB: DB,
	}
}
