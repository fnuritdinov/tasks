package service

import (
	"TaskManager/internal/models"
	"TaskManager/pkg/errors"
	"context"
	"fmt"
)

func (s *Service) CreateUser(ctx context.Context, user models.UserReq) error {
	err := user.ValidateUser()
	if err != nil {
		return err
	}

	err = s.Repo.CreateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("error from s.repo.CreateUser %w", err)
	}

	return nil
}

func (s *Service) GetUsers(ctx context.Context) ([]models.UserReq, error) {
	users, err := s.Repo.GetUsers(ctx)
	if err != nil {
		return []models.UserReq{}, fmt.Errorf("error from s.repo.GetUsers %w", err)
	}
	return users, nil
}

func (s *Service) GetUserByID(ctx context.Context, id int) (models.UserReq, error) {
	if id < 1 {
		return models.UserReq{}, errors.ErrFromValidate
	}

	user, err := s.Repo.GetUserByID(ctx, id)
	if err != nil {
		return models.UserReq{}, fmt.Errorf("error from s.repo.GetUserByID %w", err)
	}
	return user, nil
}

func (s *Service) UpdateUser(ctx context.Context, id int, updatedUser models.UserReq) error {
	if id < 1 {
		return errors.ErrFromValidate
	}

	err := updatedUser.ValidateUser()
	if err != nil {
		return err
	}

	err = s.Repo.UpdateUser(ctx, id, updatedUser)
	if err != nil {
		return fmt.Errorf("error from s.repo.UpdatedUser %w", err)
	}
	return nil
}

func (s *Service) DeleteUser(ctx context.Context, id int) error {

	err := s.Repo.DeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("error from s.repo.DeleteUser %w", err)
	}
	return nil
}
