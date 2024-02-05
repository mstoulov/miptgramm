package entity

import (
	"gorm.io/gorm"
	"time"
)

type FrontendMessage struct {
	ID         uint   `gorm:"primaryKey,autoIncrement"`
	SenderID   string `gorm:"index:chat_idx"`
	ReceiverID string `gorm:"index:chat_idx"`
	Text       string
	IsRead     bool
	Time       int64
}

type Message struct {
	FrontendMessage

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type MessageRepository interface {
	Create(*Message) error
	GetChatMessages(firstUserID string, secondUserID string) ([]Message, error)
	GetLastMessages(userID string) ([]Message, error)
	MarkReadTill(senderID string, receiverID string, time int64) (int64, error)
}

type MessageService interface {
	SendMessage(senderID, receiverID, text string) error
	GetChatMessages(firstUserID string, secondUserID string) ([]Message, error)
	GetLastMessages(userID string) ([]Message, error)
	MarkReadTill(senderID string, receiverID string, time int64) error
}
