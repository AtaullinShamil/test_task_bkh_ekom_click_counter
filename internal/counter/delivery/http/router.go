package http

import "github.com/gofiber/fiber/v2"

func MapHandlersRoutes(router fiber.Router, h *Handlers) {
	router.Get("/counter/:bannerID", h.CounterHandler())
	router.Post("/stats/:bannerID", h.StatsHandler())
}
