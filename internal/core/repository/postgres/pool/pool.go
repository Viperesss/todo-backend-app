package core_postgres_pool

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Pool interface {
}

type ConnectionPool struct {
	*pgxpool.Pool

	opTimeout time.Duration
}

func NewConnectionPool(ctx context.Context, cfg Config) (*ConnectionPool, error) {
	//  postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@app-postgres:5432/${POSTGRES_DB}?sslmode=disable
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB,
	)

	pgxConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("parse pgxconfig: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, pgxConfig)
	if err != nil {
		return nil, fmt.Errorf("create pgxpool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("pgxpool ping: %w", err)
	}

	return &ConnectionPool{
		Pool:      pool,
		opTimeout: cfg.Timeout,
	}, nil
}

func (p *ConnectionPool) OpTimeout() time.Duration {
	return p.opTimeout
}
