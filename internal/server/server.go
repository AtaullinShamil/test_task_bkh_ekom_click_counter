package server

import (
	"context"
	"fmt"

	"github.com/AtaullinShamil/test_task_bkh_ekom_click_counter/config"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	app    *fiber.App
	cfg    config.Config
	logger *logrus.Logger
	db     *sqlx.DB
	mongo  *mongo.Client
}

func NewServer(
	app *fiber.App,
	cfg config.Config,
	logger *logrus.Logger,
	db *sqlx.DB,
	mongo *mongo.Client,
) *Server {
	return &Server{
		app:    app,
		cfg:    cfg,
		logger: logger,
		db:     db,
		mongo:  mongo,
	}
}

func (s *Server) Run(ctx context.Context) error {
	s.app.Use(s.PanicRecovery())

	s.MapHandlers()

	go func() {
		if err := s.app.Listen(fmt.Sprintf("%s:%d", s.cfg.Server.Host, s.cfg.Server.Port)); err != nil {
			s.logger.Fatalf("Listen: %s", err)
		}
	}()

	<-ctx.Done()

	if !fiber.IsChild() {
		s.logger.Info("Server gracefully stopped")
	}

	return nil
}
