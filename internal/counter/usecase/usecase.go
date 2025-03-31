package usecase

import (
	"github.com/AtaullinShamil/test_task_bkh_ekom_click_counter/internal/counter"
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
