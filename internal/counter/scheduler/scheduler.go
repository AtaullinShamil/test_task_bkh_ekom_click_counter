package scheduler

import (
	"context"
	"time"

	"github.com/AtaullinShamil/test_task_bkh_ekom_click_counter/internal/counter"
	"github.com/sirupsen/logrus"
)

type Scheduler struct {
	Usecase counter.Usecase
	logger  *logrus.Logger
}

func NewScheduler(logger *logrus.Logger, useCase counter.Usecase) *Scheduler {
	return &Scheduler{
		Usecase: useCase,
		logger:  logger,
	}
}

func (s *Scheduler) Start() {
	go func() {
		now := time.Now()
		next := now.Truncate(time.Minute).Add(time.Minute)
		time.Sleep(time.Until(next))

		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()

		ctx := context.Background()

		for t := range ticker.C {
			ts := t.Truncate(time.Minute)

			s.logger.Info("TransferData started at ", ts)
			err := s.Usecase.TransferData(ctx, ts)
			if err != nil {
				s.logger.Error("TransferData err :", err)
			}
			s.logger.Info("TransferData completed successfully")
		}
	}()
}
