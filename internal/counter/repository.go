package counter

import (
	"context"
	"time"
)

type TemporaryRepository interface {
	Increase(ctx context.Context, bannerID int) error
	GetStats(ctx context.Context, bannerID int, from time.Time, to time.Time) ([]Stat, error)
}

type Repository interface {
	GetStats(ctx context.Context, bannerID int, from time.Time, to time.Time) ([]Stat, error)
}
