package main

import (
	"course_fullstack/backend/internal/handler"
	"course_fullstack/backend/internal/repository"
	repo_sqlite "course_fullstack/backend/internal/repository/sqlite"
	"course_fullstack/backend/internal/service"
	"fmt"
	"log"
	"net/http"
	"os"
)

func setupServer() *http.Server {
	dbUri, ok := os.LookupEnv("DB_URI")
	if !ok {
		log.Println("cannot get DB_URI from ENV")
		dbUri = "test.db"
	}

	db, err := repo_sqlite.NewSQLiteDB(dbUri)
	if err != nil {
		log.Panicf("Failed to initialize database: %s\n", err.Error())
	}

	repo := repository.NewRepository(db)

	service := service.NewService(repo)

	server := handler.NewServer(service)
	return server
}

func main() {
	server := setupServer()
	log.Println("server setup success")
	err := server.ListenAndServe()
	fmt.Println(err)
}
