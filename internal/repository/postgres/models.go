package postgres

import "database/sql"

type Models struct {
	Users UserModel
	Links LinkModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Users: UserModel{DB: db},
		Links: LinkModel{DB: db},
	}
}
