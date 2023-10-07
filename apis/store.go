package apis

import (
	"database/sql"

	db "github.com/aniket-mdev/hr_managment/sqlc_lib"
)

type Store struct {
	*db.Queries
	db *sql.DB
}

func NewStore(dbs *sql.DB) *Store {
	return &Store{
		db:      dbs,
		Queries: db.New(dbs),
	}
}
