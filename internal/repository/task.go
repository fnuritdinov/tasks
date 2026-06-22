package repository

import (
	"TaskManager/internal/models"
	errs "TaskManager/pkg/errors"
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (r *Repo) Create(ctx context.Context, task models.Tasks) error {
	err := r.DB.WithContext(ctx).Create(&task).Error
	if err != nil {
		return fmt.Errorf("error from db.Create %w", err)
	}

	return nil
}

func (r *Repo) Get(ctx context.Context) ([]models.Tasks, error) {
	var tasks []models.Tasks
	err := r.DB.WithContext(ctx).Find(&tasks).Error
	if err != nil {
		return []models.Tasks{}, fmt.Errorf("error from r.db.Find %w", err)
	}

	return tasks, nil

}

func (r *Repo) GetWithFilters(ctx context.Context, filter models.TaskFilter) ([]models.Tasks, error) {
	var tasks []models.Tasks

	tx := r.DB.WithContext(ctx).
		Order("id").
		Limit(filter.Limit).
		Offset(filter.Offset)
	if filter.Status != "" {
		tx = tx.Where("status = ?", filter.Status)
	}

	err := tx.Find(&tasks).Error
	if err != nil {
		return []models.Tasks{}, fmt.Errorf("error from tx.Find %w", err)
	}
	return tasks, nil
}

func (r *Repo) GetByID(ctx context.Context, id int) (models.Tasks, error) {
	var task models.Tasks
	err := r.DB.WithContext(ctx).First(&task, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Tasks{}, errs.ErrNotFound
		}
		return models.Tasks{}, fmt.Errorf("error from db.Find %w", err)
	}

	return task, nil
}

func (r *Repo) Update(ctx context.Context, id int, updatedData models.Tasks) error {

	result := r.DB.WithContext(ctx).
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

func (r *Repo) Delete(ctx context.Context, id int) error {
	result := r.DB.WithContext(ctx).Where("id = ?", id).Delete(models.Task{})
	if result.Error != nil {
		return fmt.Errorf("error from r.db.Delete %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return errs.ErrNotFound
	}

	return nil
}

func (r *Repo) Compiled(ctx context.Context, id int, updatedTask models.Tasks) (models.Tasks, error) {
	result := r.DB.WithContext(ctx).
		Model(models.Task{}).
		Where("id = ?", id).
		Update("completed", updatedTask.Completed)
	if result.Error != nil {
		return models.Tasks{}, fmt.Errorf("error from r.DB.Update %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return models.Tasks{}, errs.ErrNotFound
	}
	var task models.Tasks

	err := r.DB.WithContext(ctx).First(&task, id).Error
	if err != nil {
		return models.Tasks{}, fmt.Errorf("error from select task %w", err)
	}

	return task, nil
}

func (r *Repo) SelectTask(ctx context.Context, id int) (models.Tasks, error) {
	var task models.Tasks

	err := r.DB.WithContext(ctx).First(&task, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Tasks{}, errs.ErrNotFound
		}
		return models.Tasks{}, fmt.Errorf("error from s.DB.SelectTask %w", err)
	}
	return task, nil
}
