package models

import "time"

type Task struct {
	ID          uint      `gorm:"primaryKey"`
	Title       string    `gorm:"type:varchar(255)"`
	Description string    `gorm:"type:text"`
	Status      string    `gorm:"type:varchar(50)"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	IsDeleted   bool      `gorm:"default:false"`
}

type Tasks struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	IsDeleted   bool   `json:"IsDeleted"`
}

type TaskFilter struct {
	Status string
	Page   int
	Limit  int
	Offset int
}
