package service

import (
	"TaskManager/internal/models"
	errs "TaskManager/pkg/errors"
	"context"
	"errors"
	"fmt"
)

func (s *Service) Create(ctx context.Context, task models.Tasks) error {
	err := task.ValidateTask()
	if err != nil {
		return err
	}

	_, err = s.Repo.SelectUser(ctx, task.UserID)
	if err != nil {
		return err
	}

	err = s.Repo.Create(ctx, task)
	if err != nil {
		return fmt.Errorf("error from s.repo.Create %w", err)
	}
	return nil
}

func (s *Service) Get(ctx context.Context) ([]models.Tasks, error) {

	tasks, err := s.Repo.Get(ctx)
	if err != nil {
		return []models.Tasks{}, fmt.Errorf("error from s.repo.Get %w", err)
	}
	return tasks, nil
}

func (s *Service) GetWithFilters(ctx context.Context, filter models.TaskFilter) ([]models.Tasks, error) {

	if filter.Page <= 0 {
		filter.Page = 1
	}

	if filter.Limit <= 0 {
		filter.Limit = 10
	}

	filter.Offset = (filter.Page - 1) * filter.Limit

	tasks, err := s.Repo.GetWithFilters(ctx, filter)
	if err != nil {
		return []models.Tasks{}, fmt.Errorf("error from s.repo.Get %w", err)
	}
	return tasks, nil
}

func (s *Service) GetByID(ctx context.Context, id int) (models.Tasks, error) {
	if id < 1 {
		return models.Tasks{}, errs.ErrFromValidate
	}

	task, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return models.Tasks{}, fmt.Errorf("error from s.repo.GetByID %w", err)
	}

	return task, nil
}

func (s *Service) Update(ctx context.Context, id int, updatedTasks models.Tasks) error {
	if id < 1 {
		return errs.ErrFromValidate
	}

	err := updatedTasks.ValidateTask()
	if err != nil {
		return err
	}

	err = s.Repo.Update(ctx, id, updatedTasks)
	if err != nil {
		return fmt.Errorf("error from s.repo.Update %w", err)
	}
	return nil
}

func (s *Service) Delete(ctx context.Context, id int) error {

	err := s.Repo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("error from s.repo.Delete %w", err)
	}
	return nil
}

func (s *Service) Compiled(ctx context.Context, id int, task models.Tasks) (models.Tasks, error) {

	_, err := s.Repo.SelectTask(ctx, id)
	if err != nil {
		return models.Tasks{}, err
	}

	updatedTask, err := s.Repo.Compiled(ctx, id, task)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return models.Tasks{}, errs.ErrNotFound
		}
		return models.Tasks{}, fmt.Errorf("error from s.Repo.Complited %w", err)
	}

	return updatedTask, nil
}
