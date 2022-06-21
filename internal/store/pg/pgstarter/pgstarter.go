package pgstarter

import (
	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib" // PostgreSQL Driver
)

func NewPGStore(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
