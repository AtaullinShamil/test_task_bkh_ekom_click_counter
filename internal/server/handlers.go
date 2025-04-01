package server

import (
	"github.com/AtaullinShamil/test_task_bkh_ekom_click_counter/internal/counter/delivery/http"
	"github.com/AtaullinShamil/test_task_bkh_ekom_click_counter/internal/counter/repository"
	"github.com/AtaullinShamil/test_task_bkh_ekom_click_counter/internal/counter/scheduler"
	"github.com/AtaullinShamil/test_task_bkh_ekom_click_counter/internal/counter/usecase"
	"github.com/gofiber/swagger"
)

func (s *Server) MapHandlers() {
	s.app.Get("/swagger/*", swagger.HandlerDefault)

	posrgres := repository.NewPostgresRepository(s.db)
	mongo := repository.NewMongoRepository(s.mongo)
	usecase := usecase.NewUsecase(posrgres, mongo)
	handlers := http.NewHandlers(s.logger, usecase)
	http.MapHandlersRoutes(s.app, handlers)

	dataTransferer := scheduler.NewScheduler(s.logger, usecase)
	dataTransferer.Start()
}
