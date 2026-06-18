package main

import (
	"TaskManager/handlers"
	"TaskManager/internal/config"
	"TaskManager/internal/models"
	"TaskManager/internal/repository"
	"TaskManager/internal/service"
	db2 "TaskManager/pkg/db"
	"TaskManager/pkg/logger"
	"fmt"
	"log"
	"net/http"
)

func main() {

	lg, err := logger.New(true)
	if err != nil {
		log.Fatal("failed to create logger", err)
	}

	cfg, err := config.New("./config/config.env")
	if err != nil {
		log.Fatal("config.New", err)
	}

	db, err := db2.New(db2.Options{
		Host:     cfg.DBHOST,
		Port:     cfg.DBPORT,
		User:     cfg.DBUSER,
		Password: cfg.DBPASSWORD,
		DBName:   cfg.DBName,
	})

	err = db.AutoMigrate(&models.Task{})
	if err != nil {
		fmt.Errorf("error from db.AutoMigrate %w", err)
		return
	}

	repoTask := repository.New(db)
	serviceTask := service.New(repoTask)
	handlersTask := handlers.NewTaskHandlers(serviceTask, *lg)

	router := handlers.New(&handlersTask)

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}

}
