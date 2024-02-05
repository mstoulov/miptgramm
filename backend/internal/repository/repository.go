package repository

import (
	"course_fullstack/backend/internal/entity"
	repo_sqlite "course_fullstack/backend/internal/repository/sqlite"
	"gorm.io/gorm"
)

type Repository struct {
	User    entity.UserRepository
	Message entity.MessageRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User:    repo_sqlite.NewUserSQLite(db),
		Message: repo_sqlite.NewMessageSQLite(db),
	}
}
