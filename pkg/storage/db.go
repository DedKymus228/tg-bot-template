package storage

import (
	"context"
	"fmt"
	"tg-bot-template/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

const connStr = "postgres://%s:%s@%s:%s/%s?sslmode=disable"

type DB struct {
	pool *pgxpool.Pool
	log  *zap.Logger
}

func New(ctx context.Context, log *zap.Logger, db config.DB) (*DB, error) {
	dbURL := fmt.Sprintf(connStr, db.User, db.Password, db.Host, db.Port, db.Name)
	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatal("error create pool for DB")
	}
	err = pool.Ping(ctx)
	if err != nil {
		log.Fatal("trouble with DB pool")
	}
	return &DB{
		pool: pool,
		log:  log,
	}, nil
}

func (d *DB) ClosePool() {
	d.pool.Close()
}

func (d *DB) Ping() error {
	err := d.pool.Ping(context.Background())
	if err != nil {
		d.log.Fatal("trouble with DB pool")
	}
	return err
}
