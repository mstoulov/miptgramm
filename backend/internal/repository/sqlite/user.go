package repo_sqlite

import (
	"course_fullstack/backend/internal/entity"
	"errors"
	"gorm.io/gorm"
)

// implements entity.UserRepository interface
type UserSQLite struct {
	db *gorm.DB
}

func NewUserSQLite(db *gorm.DB) *UserSQLite {
	return &UserSQLite{db: db}
}

func (r *UserSQLite) Get(id string) (*entity.User, error) {
	var user entity.User
	if result := r.db.First(&user, "id = ?", id); result.Error == nil {
		return &user, nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, entity.ErrUserNotFound
	} else {
		return nil, result.Error
	}
}

func (r *UserSQLite) Create(user *entity.User) error {
	if result := r.db.Create(user); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			// код с вашего семинара, но не работает при попытке зарегистрировать юзера с занятым id
			return entity.ErrUserExists
		} else {
			return result.Error
		}
	} else {
		return nil
	}
}

func (r *UserSQLite) GetList(userIDs []string) ([]entity.User, error) {
	var users []entity.User
	result := r.db.Model(&entity.User{}).
		Where("id IN ?", userIDs).
		Find(&users)
	if result.Error != nil {
		return nil, result.Error
	} else {
		return users, nil
	}
}
