package models

import (
	"TaskManager/pkg/errors"
	"time"
)

type Task struct {
	ID          uint      `gorm:"primaryKey"`
	Title       string    `gorm:"type:varchar(255)"`
	Description string    `gorm:"type:text"`
	Status      string    `gorm:"type:varchar(50)"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	IsDeleted   bool      `gorm:"default:false"`
	UserID      int       `gorm:"integer"`
	Completed   bool      `gorm:"default:false"`
}

type User struct {
	ID        uint      `gorm:"primeryKey"`
	Name      string    `gorm:"type:varchar(50)"`
	Age       int       `gorm:"integer"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type UserReq struct {
	ID   int
	Name string
	Age  int
}

func (u *UserReq) ValidateUser() error {
	if len(u.Name) == 0 {
		return errors.ErrFromValidate
	}
	if u.Age < 1 {
		return errors.ErrFromValidate
	}
	return nil
}

type Tasks struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   time.Time
	IsDeleted   bool `json:"IsDeleted"`
	UserID      int  `json:"userID"`
	Completed   bool `json:"completed"`
}

type TaskFilter struct {
	Status string
	Page   int
	Limit  int
	Offset int
}

func (t *Tasks) ValidateTask() error {
	if len(t.Title) == 0 && len(t.Description) == 0 && len(t.Status) == 0 {
		return errors.ErrFromValidate
	}
	return nil
}
