package repositories

import "ddd-go/internal/domain/entities"

// UserRepository defines the methods for user data storage.
type UserRepository interface {
	CreateUser(user *entities.User) error
	GetUserByID(id string) (*entities.User, error)
	GetAllUsers() ([]*entities.User, error)
}
