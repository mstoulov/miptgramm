package repo_sqlite

import (
	"course_fullstack/backend/internal/entity"
	"gorm.io/gorm"
	"slices"
)

// implements entity.UserRepository interface
type MessageSQLite struct {
	db *gorm.DB
}

func NewMessageSQLite(db *gorm.DB) *MessageSQLite {
	return &MessageSQLite{db: db}
}

func (r *MessageSQLite) Create(message *entity.Message) error {
	result := r.db.Create(message)
	return result.Error
}

func (r *MessageSQLite) GetChatMessages(firstUserID string, secondUserID string) ([]entity.Message, error) {
	var allMessages []entity.Message
	result := r.db.Model(&entity.Message{}).
		Where("(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)",
			firstUserID, secondUserID, secondUserID, firstUserID).
		Find(&allMessages)
	if result.Error != nil {
		return nil, result.Error
	}
	return allMessages, nil
}

func (r *MessageSQLite) GetLastMessages(userID string) ([]entity.Message, error) {
	var allMessages []entity.Message
	result := r.db.Model(&entity.Message{}).
		Where("sender_id = ? OR receiver_id = ?", userID, userID).
		Find(&allMessages)
	if result.Error != nil {
		return nil, result.Error
	}
	var lastMessages []entity.Message
	tmp := map[string][]entity.Message{}
	for _, message := range allMessages {
		var chatID string
		if message.SenderID == userID {
			chatID = message.ReceiverID
		} else {
			chatID = message.SenderID
		}
		tmp[chatID] = append(tmp[chatID], message)
	}

	msgCmp := func(message1 entity.Message, message2 entity.Message) int {
		if message1.Time < message2.Time {
			return 1
		} else if message1.Time > message2.Time {
			return -1
		} else {
			return 0
		}
	}

	for _, messages := range tmp {
		lastMessages = append(lastMessages, slices.MinFunc(messages, msgCmp))
	}
	slices.SortFunc(lastMessages, msgCmp)
	return lastMessages, nil
}

func (r *MessageSQLite) MarkReadTill(senderID string, receiverID string, time int64) (int64, error) {
	result := r.db.Model(&entity.Message{}).
		Where("sender_id = ? AND receiver_id = ? AND time <= ? AND is_read = false", senderID, receiverID, time).
		Update("is_read", true)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
