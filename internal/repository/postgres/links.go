package postgres

import (
	"database/sql"

	"github.com/binsabit/urlshortener/internal/repository"
)

type LinkModel struct {
	DB *sql.DB
}

// type LinkStorage interface {
// 	SaveURL(int64, Link) (int64, error)
// 	DeleteURL(int64) error
// 	UpdateURL(Link)
// }

func (l LinkModel) SaveURL(userID int64, link repository.Link) error {
	return nil
}

func (l LinkModel) DeleteURL(linkID int64) error {
	return nil
}

func (l LinkModel) UpdateURL(link repository.Link) error {
	return nil
}
