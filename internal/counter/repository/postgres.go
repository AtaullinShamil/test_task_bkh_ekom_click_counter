package repository

import (
	"context"
	"time"

	"github.com/AtaullinShamil/test_task_bkh_ekom_click_counter/internal/counter"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type PostgresRepository struct {
	db *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (p *PostgresRepository) GetStats(ctx context.Context, bannerID int, from time.Time, to time.Time) ([]counter.Stat, error) {
	var stats []counter.Stat

	err := p.db.SelectContext(context.Background(),
		&stats,
		"SELECT timestamp, count FROM counter.clicks WHERE banner_id=$1 AND timestamp BETWEEN $2 AND $3",
		bannerID, from, to)
	if err != nil {
		return nil, errors.Wrap(err, "SelectContext")
	}

	return stats, nil
}
