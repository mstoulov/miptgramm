package entity

import (
	"gorm.io/gorm"
	"time"
)

type FrontendUser struct {
	ID   string `gorm:"primaryKey"`
	Name string `gorm:"not null"`
}

type User struct {
	FrontendUser
	Password string `gorm:"not null" json:"password,omitempty"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type UserRepository interface {
	Create(*User) error
	Get(id string) (*User, error)
	GetList(ids []string) ([]User, error)
}

type UserService interface {
	Register(id string, name string, password string) error
	Login(id string, password string, client Client) error
	Logout(id string)
	GetUser(id string) (*User, error)
	GetUserList(ids []string) ([]User, error)
}
