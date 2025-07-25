package db

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres() (*Postgres, error) {
	db, err := sql.Open("postgres", "postgres://admin:admin@localhost:5432/techcrm?sslmode=disable")

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("Connected to PostgreSQL")

	return &Postgres{db: db}, nil
}

func (p *Postgres) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return p.db.ExecContext(ctx, query, args...)
}

func (p *Postgres) Close() error {
	return p.db.Close()
}
