package service

import (
	"TaskManager/internal/models"
	"TaskManager/internal/repository"
	"TaskManager/pkg/errors"
	"context"
	"fmt"
)

type TaskService interface {
	Create(ctx context.Context, task models.Task) error
	Get(ctx context.Context) ([]models.Task, error)
	GetWithFilters(ctx context.Context, filter models.TaskFilter) ([]models.Task, error)
	GetByID(ctx context.Context, id int) (models.Task, error)
	Update(ctx context.Context, id int, updatedTasks models.Tasks) error
	Delete(ctx context.Context, id int) error
}

type service struct {
	repo repository.TaskRepo
}

func New(repo repository.TaskRepo) TaskService {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(ctx context.Context, task models.Task) error {
	if len(task.Description) == 0 {
		return errors.ErrFromValidate
	}

	if len(task.Title) == 0 {
		return errors.ErrFromValidate
	}

	if len(task.Status) == 0 {
		return errors.ErrFromValidate
	}

	err := s.repo.Create(ctx, task)
	if err != nil {
		return fmt.Errorf("error from s.repo.Create %w", err)
	}
	return nil
}

func (s *service) Get(ctx context.Context) ([]models.Task, error) {

	tasks, err := s.repo.Get(ctx)
	if err != nil {
		return []models.Task{}, fmt.Errorf("error from s.repo.Get %w", err)
	}
	return tasks, nil
}

func (s *service) GetWithFilters(ctx context.Context, filter models.TaskFilter) ([]models.Task, error) {

	if filter.Page <= 0 {
		filter.Page = 1
	}

	if filter.Limit <= 0 {
		filter.Limit = 10
	}

	filter.Offset = (filter.Page - 1) * filter.Limit

	tasks, err := s.repo.GetWithFilters(ctx, filter)
	if err != nil {
		return []models.Task{}, fmt.Errorf("error from s.repo.Get %w", err)
	}
	return tasks, nil
}

func (s *service) GetByID(ctx context.Context, id int) (models.Task, error) {
	if id < 1 {
		return models.Task{}, errors.ErrFromValidate
	}

	task, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return models.Task{}, fmt.Errorf("error from s.repo.GetByID %w", err)
	}

	return task, nil
}

func (s *service) Update(ctx context.Context, id int, updatedTasks models.Tasks) error {
	if id < 1 {
		return errors.ErrFromValidate
	}

	if len(updatedTasks.Title) == 0 && len(updatedTasks.Status) == 0 && len(updatedTasks.Description) == 0 {
		return errors.ErrFromValidate
	}

	err := s.repo.Update(ctx, id, updatedTasks)
	if err != nil {
		return fmt.Errorf("error from s.repo.Update %w", err)
	}
	return nil
}

func (s *service) Delete(ctx context.Context, id int) error {

	err := s.repo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("error from s.repo.Delete %w", err)
	}
	return nil
}
