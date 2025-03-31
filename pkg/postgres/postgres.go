package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

func NewPostgresClient(ctx context.Context, postgresConfig Config) (*sqlx.DB, error) {
	cfg, err := pgx.ParseConfig(postgresConfig.DSN)
	if err != nil {
		return nil, errors.Wrap(err, "ParseConfig")
	}

	cfg.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	db := sqlx.NewDb(stdlib.OpenDB(*cfg, nil...), "pgx")
	if err = db.PingContext(ctx); err != nil {
		return nil, errors.Wrap(err, "db.PingContext")
	}

	db.SetMaxOpenConns(postgresConfig.MaxOpenConns)
	db.SetMaxIdleConns(postgresConfig.MaxIdleConns)

	db.SetConnMaxLifetime(postgresConfig.ConnMaxLifetime)
	db.SetConnMaxIdleTime(postgresConfig.ConnMaxIdleTime)

	return db, nil
}
