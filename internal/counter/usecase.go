package counter

import (
	"context"
	"time"
)

type Usecase interface {
	Counter(ctx context.Context, bannerID int) error
	Stats(ctx context.Context, bannerID int, ts StatsRequest) (*StatsResponse, error)
	TransferData(ctx context.Context, periodEnd time.Time) error
}
