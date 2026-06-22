package repository

import (
	"TaskManager/internal/models"
	errs "TaskManager/pkg/errors"
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (r *Repo) CreateUser(ctx context.Context, user models.UserReq) error {
	err := r.DB.WithContext(ctx).Create(&user).Error
	if err != nil {
		return fmt.Errorf("error from r.db.Create %w", err)
	}
	return nil
}

func (r *Repo) GetUsers(ctx context.Context) ([]models.UserReq, error) {
	var users []models.UserReq

	err := r.DB.WithContext(ctx).Find(&users).Error
	if err != nil {
		return []models.UserReq{}, fmt.Errorf("error from r.db.Find %w", err)
	}
	return users, nil
}

func (r *Repo) GetUserByID(ctx context.Context, id int) (models.UserReq, error) {
	var user models.UserReq

	err := r.DB.WithContext(ctx).First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.UserReq{}, errs.ErrNotFound
		}
		return models.UserReq{}, fmt.Errorf("error from r.db.GetUserByID %w", err)
	}
	return user, nil
}

func (r *Repo) UpdateUser(ctx context.Context, id int, updatedUser models.UserReq) error {
	result := r.DB.WithContext(ctx).Model(models.UserReq{}).
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

func (r *Repo) DeleteUser(ctx context.Context, id int) error {
	result := r.DB.WithContext(ctx).
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

func (r *Repo) SelectUser(ctx context.Context, userID int) (models.User, error) {
	var user models.User

	err := r.DB.WithContext(ctx).First(&user, userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, errs.ErrNotFound
		}
		return models.User{}, fmt.Errorf("error from r.db.SelectUser %w", err)
	}
	return user, nil
}
