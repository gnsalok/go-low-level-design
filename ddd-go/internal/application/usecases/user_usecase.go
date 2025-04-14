package usecases

import (
	"ddd-go/internal/domain/entities"
	"ddd-go/internal/domain/repositories"
	"errors"
)

type UserUseCase struct {
	userRepo repositories.UserRepository
}

func NewUserUseCase(userRepo repositories.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

func (uc *UserUseCase) RegisterUser(name string, email string) (*entities.User, error) {
	if name == "" || email == "" {
		return nil, errors.New("name and email cannot be empty")
	}

	user := &entities.User{
		Name:  name,
		Email: email,
	}

	err := uc.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserUseCase) GetUserByID(id string) (*entities.User, error) {
	user, err := uc.userRepo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserUseCase) GetAllUsers() ([]*entities.User, error) {
	users, err := uc.userRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}
