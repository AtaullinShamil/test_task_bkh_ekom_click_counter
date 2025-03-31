package http

import (
	"github.com/AtaullinShamil/test_task_bkh_ekom_click_counter/internal/counter"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Handlers struct {
	Usecase counter.Usecase
	logger  *logrus.Logger
}

func NewHandlers(logger *logrus.Logger, useCase counter.Usecase) *Handlers {
	return &Handlers{
		Usecase: useCase,
		logger:  logger,
	}
}

func (h *Handlers) incrementCounter() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	}
}

func (h *Handlers) getStats() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	}
}
