// filepath: /internal/infrastructure/database/user_repository_sql.go
package database

import (
	"database/sql"
	"ddd-go/internal/domain/entities"
	"ddd-go/internal/domain/repositories"
)

type UserRepositorySQL struct {
	db *sql.DB
}

func NewUserRepositorySQL(db *sql.DB) repositories.UserRepository {
	return &UserRepositorySQL{db: db}
}

func (r *UserRepositorySQL) CreateUser(user *entities.User) error {
	_, err := r.db.Exec("INSERT INTO users (id, name, email) VALUES (?, ?, ?)", user.ID, user.Name, user.Email)
	return err
}

func (r *UserRepositorySQL) GetUserByID(id string) (*entities.User, error) {
	row := r.db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id)
	user := &entities.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositorySQL) GetAllUsers() ([]*entities.User, error) {
	rows, err := r.db.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entities.User
	for rows.Next() {
		user := &entities.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
