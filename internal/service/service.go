package service

import (
	"TaskManager/internal/models"
	"TaskManager/internal/repository"
	"TaskManager/pkg/errors"
	"context"
	"fmt"
)

type TaskService interface {
	Create(ctx context.Context, task models.Tasks) error
	Get(ctx context.Context) ([]models.Task, error)
	GetWithFilters(ctx context.Context, filter models.TaskFilter) ([]models.Task, error)
	GetByID(ctx context.Context, id int) (models.Task, error)
	Update(ctx context.Context, id int, updatedTasks models.Tasks) error
	Delete(ctx context.Context, id int) error
	CreateUser(ctx context.Context, user models.UserReq) error
	GetUsers(ctx context.Context) ([]models.UserReq, error)
	GetUserByID(ctx context.Context, id int) (models.UserReq, error)
	UpdateUser(ctx context.Context, id int, updatedUser models.UserReq) error
	DeleteUser(ctx context.Context, id int) error
}

type service struct {
	repo repository.TaskRepo
}

func New(repo repository.TaskRepo) TaskService {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(ctx context.Context, task models.Tasks) error {
	err := task.ValidateTask()
	if err != nil {
		return err
	}

	_, err = s.repo.SelectUser(ctx, task.UserID)
	if err != nil {
		return err
	}

	err = s.repo.Create(ctx, task)
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

	err := updatedTasks.ValidateTask()
	if err != nil {
		return err
	}

	err = s.repo.Update(ctx, id, updatedTasks)
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

func (s *service) CreateUser(ctx context.Context, user models.UserReq) error {
	err := user.ValidateUser()
	if err != nil {
		return err
	}

	err = s.repo.CreateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("error from s.repo.CreateUser %w", err)
	}

	return nil
}

func (s *service) GetUsers(ctx context.Context) ([]models.UserReq, error) {
	users, err := s.repo.GetUsers(ctx)
	if err != nil {
		return []models.UserReq{}, fmt.Errorf("error from s.repo.GetUsers %w", err)
	}
	return users, nil
}

func (s *service) GetUserByID(ctx context.Context, id int) (models.UserReq, error) {
	if id < 1 {
		return models.UserReq{}, errors.ErrFromValidate
	}

	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return models.UserReq{}, fmt.Errorf("error from s.repo.GetUserByID %w", err)
	}
	return user, nil
}

func (s *service) UpdateUser(ctx context.Context, id int, updatedUser models.UserReq) error {
	if id < 1 {
		return errors.ErrFromValidate
	}

	err := updatedUser.ValidateUser()
	if err != nil {
		return err
	}

	err = s.repo.UpdateUser(ctx, id, updatedUser)
	if err != nil {
		fmt.Errorf("error from s.repo.UpdatedUser %w", err)
	}
	return nil
}

func (s *service) DeleteUser(ctx context.Context, id int) error {

	err := s.repo.DeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("error from s.repo.DeleteUser")
	}
	return nil
}
