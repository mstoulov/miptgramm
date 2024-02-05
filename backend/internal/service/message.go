package service

import (
	"course_fullstack/backend/internal/entity"
	"time"
)

func (s *Service) GetChatMessages(firstUserID string, secondUserID string) ([]entity.Message, error) {
	return s.messageRepo.GetChatMessages(firstUserID, secondUserID)
}

func (s *Service) GetLastMessages(userID string) ([]entity.Message, error) {
	return s.messageRepo.GetLastMessages(userID)
}

func (s *Service) MarkReadTill(senderID string, receiverID string, time int64) error {
	markedCnt, err := s.messageRepo.MarkReadTill(senderID, receiverID, time)
	if err != nil {
		return err
	}
	if markedCnt > 0 {
		s.pingUser(senderID)
		s.pingUser(receiverID)
	}
	return nil
}

func (s *Service) SendMessage(senderID, receiverID, text string) error {
	err := s.messageRepo.Create(&entity.Message{
		FrontendMessage: entity.FrontendMessage{
			SenderID:   senderID,
			ReceiverID: receiverID,
			Text:       text,
			IsRead:     false,
			Time:       time.Now().UnixMilli(),
		},
	})
	if err != nil {
		return err
	}
	s.pingUser(receiverID)
	s.pingUser(senderID)
	return nil
}
