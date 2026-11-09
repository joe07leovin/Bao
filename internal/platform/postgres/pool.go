package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joe07leovin/Bao/internal/platform/config"
)

func Open(ctx context.Context, c config.Config) (*pgxpool.Pool, error) {
	p, err := pgxpool.ParseConfig(c.DatabaseURL)
	if err != nil {
		return nil, errors.New("invalid database connection configuration")
	}
	if p.ConnConfig.TLSConfig == nil || p.ConnConfig.TLSConfig.InsecureSkipVerify || len(p.ConnConfig.Fallbacks) > 0 {
		return nil, errors.New("database requires verified TLS without fallback")
	}
	p.MaxConns = c.DatabaseMaxConnections
	p.ConnConfig.ConnectTimeout = c.DatabaseConnectTimeout
	pool, err := pgxpool.NewWithConfig(ctx, p)
	if err != nil {
		return nil, errors.New("database pool creation failed")
	}
	check, cancel := context.WithTimeout(ctx, c.DatabaseConnectTimeout)
	defer cancel()
	if err := pool.Ping(check); err != nil {
		pool.Close()
		return nil, errors.New("database connection failed")
	}
	return pool, nil
}
