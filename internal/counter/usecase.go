package counter

import "context"

type Usecase interface {
	Counter(ctx context.Context, bannerID int) error
	Stats(ctx context.Context, bannerID int, ts StatsRequest) (*StatsResponse, error)
}
