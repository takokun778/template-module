package database

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type Database struct {
	Client *sql.DB
}

func New(
	url string,
) (*Database, error) {
	db, err := sql.Open("pgx", url)
	if err != nil {
		return nil, fmt.Errorf("failed to new database: %w", err)
	}

	return &Database{
		Client: db,
	}, nil
}
