package services

import (
	"ddd-go/internal/domain/entities"
	"ddd-go/internal/domain/repositories"
	"errors"
)

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(name, email string) (*entities.User, error) {
	if name == "" || email == "" {
		return nil, errors.New("name and email are required")
	}

	user := &entities.User{
		Name:  name,
		Email: email,
	}

	err := s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) AuthenticateUser(email string) (*entities.User, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}
