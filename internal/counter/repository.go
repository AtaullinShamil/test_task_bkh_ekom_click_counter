package counter

import (
	"context"
	"time"
)

type TemporaryRepository interface {
	Increase(ctx context.Context, bannerID int) error
	GetStatsBeforeTime(ctx context.Context, ts time.Time) ([]Stat, error)
	DeleteStatsBeforeTime(ctx context.Context, ts time.Time) error
}

type Repository interface {
	GetStats(ctx context.Context, bannerID int, from time.Time, to time.Time) ([]Stat, error)
	SaveStats(ctx context.Context, stats []Stat) error
}
