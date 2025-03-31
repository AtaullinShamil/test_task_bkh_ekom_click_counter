package http

import (
	"net/http"

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

func (h *Handlers) CounterHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.Context()

		bannerID, err := c.ParamsInt("bannerID")
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("Invalid banner ID")
		}
		err = h.Usecase.Counter(ctx, bannerID)
		if err != nil {
			h.logger.Errorf("Counter err: %v", err)
			return c.Status(fiber.StatusBadRequest).JSON(counter.ErrorResponse{
				Error: err.Error(),
			})
		}

		return c.SendString("OK")
	}
}

func (h *Handlers) StatsHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.Context()

		bannerID, err := c.ParamsInt("bannerID")
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("Invalid banner ID")
		}

		var req counter.StatsRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
		}

		resp, err := h.Usecase.Stats(ctx, bannerID, req)
		if err != nil {
			h.logger.Errorf("Stats err: %v", err)
			return c.Status(fiber.StatusBadRequest).JSON(counter.ErrorResponse{
				Error: err.Error(),
			})
		}

		return c.JSON(resp)
	}
}
