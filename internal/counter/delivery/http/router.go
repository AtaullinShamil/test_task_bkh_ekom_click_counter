package http

import "github.com/gofiber/fiber/v2"

func MapHandlersRoutes(router fiber.Router, h *Handlers) {
	router.Get("/counter/:bannerID", h.incrementCounter())
	router.Post("/stats/:bannerID", h.getStats())
}
