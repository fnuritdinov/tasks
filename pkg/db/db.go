package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Options struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func New(o Options) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		o.Host,
		o.User,
		o.Password,
		o.DBName,
		o.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return &gorm.DB{}, fmt.Errorf("error from gorm.Open")
	}

	return db, nil
}
