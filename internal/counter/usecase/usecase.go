package usecase

import (
	"context"

	"github.com/AtaullinShamil/test_task_bkh_ekom_click_counter/internal/counter"
	"github.com/pkg/errors"
)

type Usecase struct {
	repository          counter.Repository
	temporaryRepository counter.TemporaryRepository
}

func NewUsecase(repository counter.Repository,
	temporaryRepository counter.TemporaryRepository) *Usecase {
	return &Usecase{
		repository:          repository,
		temporaryRepository: temporaryRepository,
	}
}

func (u *Usecase) Counter(ctx context.Context, bannerID int) error {
	err := u.temporaryRepository.Increase(ctx, bannerID)
	if err != nil {
		return errors.Wrap(err, "Increase")
	}

	return nil
}

func (u *Usecase) Stats(ctx context.Context, bannerId int, ts counter.StatsRequest) (*counter.StatsResponse, error) {
	var stats []counter.Stat
	var err error

	stats, err = u.temporaryRepository.GetStats(ctx, bannerId, ts.From, ts.To)
	if err != nil {
		return nil, errors.Wrap(err, "temporaryRepository.GetStats")
	}

	if len(stats) == 0 {
		stats, err = u.repository.GetStats(ctx, bannerId, ts.From, ts.To)
		if err != nil {
			return nil, errors.Wrap(err, "repository.GetStats")
		}
	}

	return &counter.StatsResponse{
		Stats: stats,
	}, nil
}
