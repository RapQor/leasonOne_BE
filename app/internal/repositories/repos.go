package repositories

import "database/sql"

type repos struct {
	db *sql.DB
}

func Repos(db *sql.DB) *repos {
	return &repos{db: db}
}
