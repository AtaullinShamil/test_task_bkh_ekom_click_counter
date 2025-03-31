package server

import "github.com/gofiber/fiber/v2"

func (s *Server) PanicRecovery() fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				s.logger.Errorf("Recovered from panic: %v\n", r)
				c.Status(fiber.StatusInternalServerError)
			}
		}()

		return c.Next()
	}
}
