package service

import "course_fullstack/backend/internal/entity"

func (s *Service) Register(id, name, password string, client entity.Client) error {
	if id == "" || len(id) > 32 {
		return entity.ErrInvalidUserID
	}
	if name == "" || len(name) > 32 {
		return entity.ErrInvalidUserName
	}
	if password == "" || len(password) > 32 {
		return entity.ErrInvalidPassword
	}
	err := s.userRepo.Create(&entity.User{
		FrontendUser: entity.FrontendUser{
			ID:   id,
			Name: name,
		},
		Password: password,
	})
	if err != nil {
		return err
	}
	return s.Login(id, password, client)
}

func (s *Service) Login(id string, password string, client entity.Client) error {
	user, err := s.GetUser(id)
	if err != nil {
		return err
	}
	if user.Password != password {
		return entity.ErrIncorrectPassword
	}
	s.addClient(id, client)
	return nil
}

func (s *Service) Logout(id string) {
	s.removeClient(id)
}

func (s *Service) GetUser(id string) (*entity.User, error) {
	return s.userRepo.Get(id)
}

func (s *Service) GetUserList(ids []string) ([]entity.User, error) {
	return s.userRepo.GetList(ids)
}
