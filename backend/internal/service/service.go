package service

import (
	"course_fullstack/backend/internal/entity"
	"course_fullstack/backend/internal/repository"
	"sync"
)

type Service struct {
	userRepo    entity.UserRepository
	messageRepo entity.MessageRepository

	clientConnectionMutex sync.RWMutex
	clientConnections     map[string]entity.Client
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		userRepo:          repo.User,
		messageRepo:       repo.Message,
		clientConnections: make(map[string]entity.Client),
	}
}

func (s *Service) pingUser(id string) {
	s.clientConnectionMutex.RLock()
	client, ok := s.clientConnections[id]
	s.clientConnectionMutex.RUnlock()
	if ok {
		client.Ping()
	}
}

func (s *Service) addClient(id string, client entity.Client) {
	s.clientConnectionMutex.Lock()
	s.clientConnections[id] = client
	s.clientConnectionMutex.Unlock()
}

func (s *Service) removeClient(id string) {
	s.clientConnectionMutex.Lock()
	delete(s.clientConnections, id)
	s.clientConnectionMutex.Unlock()
}
