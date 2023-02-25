package postgres

import (
	"database/sql"

	"github.com/binsabit/urlshortener/internal/repository"
)

type UserModel struct {
	DB *sql.DB
}

// type UserStorage interface {
// 	CreateUser(User) (int64, error)
// 	DeleteUser(int64) error
// 	GetUserByID(int64) (User, error)
// 	UpdateUser(int64, User) (User, int64)
// }
func (u UserModel) CreateUser(user *repository.User) (int64, error) {
	return 0, nil
}

func (u UserModel) DeleteUser(id int64) error {
	return nil
}

func (u UserModel) GetUserByID(id int64) (*repository.User, error) {
	return nil, nil
}

func (u UserModel) UpdateUser(id int64, user repository.User) error {
	return nil
}
