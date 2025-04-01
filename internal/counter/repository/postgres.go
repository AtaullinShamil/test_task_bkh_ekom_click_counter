package repository

import (
	"context"
	"fmt"
	"strings"
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

	err := p.db.SelectContext(ctx,
		&stats,
		"SELECT timestamp, count FROM counter.clicks WHERE banner_id=$1 AND timestamp BETWEEN $2 AND $3",
		bannerID, from, to)
	if err != nil {
		return nil, errors.Wrap(err, "SelectContext")
	}

	return stats, nil
}

func (p *PostgresRepository) SaveStats(ctx context.Context, stats []counter.Stat) error {
	if len(stats) == 0 {
		return nil
	}

	query := "INSERT INTO counter.clicks (timestamp, count, banner_id) VALUES "
	values := []interface{}{}
	placeholders := []string{}

	for i, stat := range stats {
		idx := i * 3
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d, $%d)", idx+1, idx+2, idx+3))
		values = append(values, stat.Timestamp, stat.Value, stat.BannerID)
	}

	query += strings.Join(placeholders, ", ")
	_, err := p.db.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "ExecContext")
	}

	return nil
}
