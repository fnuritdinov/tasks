package repository

import (
	"TaskManager/internal/models"
	errs "TaskManager/pkg/errors"
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type TaskRepo interface {
	Create(ctx context.Context, task models.Tasks) error
	Get(ctx context.Context) ([]models.Task, error)
	GetWithFilters(ctx context.Context, filter models.TaskFilter) ([]models.Task, error)
	GetByID(ctx context.Context, id int) (models.Task, error)
	Update(ctx context.Context, id int, updatedData models.Tasks) error
	Delete(ctx context.Context, id int) error
	CreateUser(ctx context.Context, user models.UserReq) error
	GetUsers(ctx context.Context) ([]models.UserReq, error)
	GetUserByID(ctx context.Context, id int) (models.UserReq, error)
	UpdateUser(ctx context.Context, id int, updatedUser models.UserReq) error
	DeleteUser(ctx context.Context, id int) error
	SelectUser(ctx context.Context, userID int) (models.Tasks, error)
}

type repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) TaskRepo {
	return &repo{
		db: db,
	}
}

func (r *repo) Create(ctx context.Context, task models.Tasks) error {
	err := r.db.WithContext(ctx).Create(&task).Error
	if err != nil {
		return fmt.Errorf("error from db.Create %w", err)
	}

	return nil
}

func (r *repo) Get(ctx context.Context) ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.WithContext(ctx).Find(&tasks).Error
	if err != nil {
		return []models.Task{}, fmt.Errorf("error from r.db.Find %w", err)
	}

	return tasks, nil

}

func (r *repo) GetWithFilters(ctx context.Context, filter models.TaskFilter) ([]models.Task, error) {
	var tasks []models.Task

	tx := r.db.WithContext(ctx).
		Order("id").
		Limit(filter.Limit).
		Offset(filter.Offset)
	if filter.Status != "" {
		tx = tx.Where("status = ?", filter.Status)
	}

	err := tx.Find(&tasks).Error
	if err != nil {
		return []models.Task{}, fmt.Errorf("error from tx.Find %w", err)
	}
	return tasks, nil
}

func (r *repo) GetByID(ctx context.Context, id int) (models.Task, error) {
	var task models.Task
	err := r.db.WithContext(ctx).First(&task, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Task{}, errs.ErrNotFound
		}
		return models.Task{}, fmt.Errorf("error from db.Find %w", err)
	}

	return task, nil
}

func (r *repo) Update(ctx context.Context, id int, updatedData models.Tasks) error {

	result := r.db.WithContext(ctx).
		Model(&models.Task{}).
		Where("id = ?", id).
		Update("status", updatedData.Status).
		Update("title", updatedData.Title).
		Update("description", updatedData.Description)

	if result.Error != nil {
		return fmt.Errorf("error from r.db.Update %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return errs.ErrNotFound
	}

	return nil
}

func (r *repo) Delete(ctx context.Context, id int) error {
	result := r.db.WithContext(ctx).Where("id = ?", id).Delete(models.Task{})
	if result.Error != nil {
		return fmt.Errorf("error from r.db.Delete %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return errs.ErrNotFound
	}

	return nil
}

func (r *repo) CreateUser(ctx context.Context, user models.UserReq) error {
	err := r.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return fmt.Errorf("error from r.db.Create %w", err)
	}
	return nil
}

func (r *repo) GetUsers(ctx context.Context) ([]models.UserReq, error) {
	var users []models.UserReq

	err := r.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return []models.UserReq{}, fmt.Errorf("error from r.db.Find %w", err)
	}
	return users, nil
}

func (r *repo) GetUserByID(ctx context.Context, id int) (models.UserReq, error) {
	var user models.UserReq

	err := r.db.WithContext(ctx).First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.UserReq{}, errs.ErrNotFound
		}
		return models.UserReq{}, fmt.Errorf("error from r.db.GetUserByID %w", err)
	}
	return user, nil
}

func (r *repo) UpdateUser(ctx context.Context, id int, updatedUser models.UserReq) error {
	result := r.db.WithContext(ctx).Model(models.UserReq{}).
		Where("id = ?", id).
		Update("name", updatedUser.Name).
		Update("age", updatedUser.Age)
	if result.Error != nil {
		return fmt.Errorf("error from r.db.Update %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return errs.ErrNotFound
	}
	return nil
}

func (r *repo) DeleteUser(ctx context.Context, id int) error {
	result := r.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(models.UserReq{})
	if result.Error != nil {
		return fmt.Errorf("error from r.db.Delete %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return errs.ErrNotFound
	}
	return nil
}

func (r *repo) SelectUser(ctx context.Context, userID int) (models.Tasks, error) {
	var user models.Tasks

	err := r.db.WithContext(ctx).First(&user, userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Tasks{}, errs.ErrNotFound
		}
		return models.Tasks{}, fmt.Errorf("error from r.db.SelectUser %w", err)
	}
	return user, nil
}
