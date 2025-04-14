// filepath: /internal/infrastructure/database/user_repository_memory.go
package database

import (
	"ddd-go/internal/domain/entities"
	"ddd-go/internal/domain/repositories"
	"errors"
)

type UserRepositoryMemory struct {
	db map[string]*entities.User
}

func NewUserRepositoryMemory() repositories.UserRepository {
	return &UserRepositoryMemory{
		db: make(map[string]*entities.User),
	}
}

func (r *UserRepositoryMemory) CreateUser(user *entities.User) error {
	if _, exists := r.db[user.ID]; exists {
		return errors.New("user already exists")
	}
	r.db[user.ID] = user
	return nil
}

func (r *UserRepositoryMemory) GetUserByID(id string) (*entities.User, error) {
	user, exists := r.db[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *UserRepositoryMemory) GetAllUsers() ([]*entities.User, error) {
	users := []*entities.User{}
	for _, user := range r.db {
		users = append(users, user)
	}
	return users, nil
}
