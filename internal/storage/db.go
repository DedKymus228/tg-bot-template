package storage

import (
	"context"
	"fmt"
	"log"
	"tg-bot-template/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

const connStr = "postgres://%s:%s@%s:%s/%s?sslmode=disable"

type DB struct {
	pool *pgxpool.Pool
}

func New(ctx context.Context, db config.DB) (*DB, error) {
	dbURL := fmt.Sprintf(connStr, db.User, db.Password, db.Host, db.Port, db.Name)
	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatalf("error create pool for DB")
		return nil, err
	}
	err = pool.Ping(ctx)
	if err != nil {
		log.Fatalf("trouble with DB pool")
		return nil, err
	}
	return &DB{
		pool: pool,
	}, nil
}

func (d *DB) ClosePool() {
	d.pool.Close()
}
