package usecase

import (
	"context"
	"time"

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
	from, to, err := validateTimestamp(ts.From, ts.To)
	if err != nil {
		return nil, errors.Wrap(err, "validateTimestamp")
	}

	stats, err := u.repository.GetStats(ctx, bannerId, from, to)
	if err != nil {
		return nil, errors.Wrap(err, "repository.GetStats")
	}

	return &counter.StatsResponse{
		Stats: stats,
	}, nil
}

func (u *Usecase) TransferData(ctx context.Context, periodEnd time.Time) error {
	stats, err := u.temporaryRepository.GetStatsBeforeTime(ctx, periodEnd)
	if err != nil {
		return errors.Wrap(err, "GetStatsBeforeTime")
	}

	err = u.repository.SaveStats(ctx, stats)
	if err != nil {
		return errors.Wrap(err, "SaveStats")
	}

	err = u.temporaryRepository.DeleteStatsBeforeTime(ctx, periodEnd)
	if err != nil {
		return errors.Wrap(err, "DeleteStatsBeforeTime")
	}

	return nil
}
