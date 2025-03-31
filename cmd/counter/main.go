package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/AtaullinShamil/test_task_bkh_ekom_click_counter/config"
	"github.com/AtaullinShamil/test_task_bkh_ekom_click_counter/internal/server"
	"github.com/AtaullinShamil/test_task_bkh_ekom_click_counter/pkg/mongo"
	"github.com/AtaullinShamil/test_task_bkh_ekom_click_counter/pkg/postgres"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)
	defer stop()

	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	cfg, err := config.Load()
	if err != nil {
		logger.Fatalf("Load: %s", err)
	}

	db, err := postgres.NewPostgresClient(ctx, cfg.Postgres)
	if err != nil {
		logger.Fatalf("NewPostgresClient: %s", err)
	}
	defer db.Close()

	mongo, err := mongo.NewMongoClient(ctx, cfg.Mongo)
	if err != nil {
		logger.Fatalf("NewMongoClient: %s", err)
	}
	defer mongo.Disconnect(ctx)

	app := fiber.New(fiber.Config{
		Prefork:     false,
		JSONEncoder: jsoniter.Marshal,
		JSONDecoder: jsoniter.Unmarshal,
	})
	defer app.Shutdown()

	s := server.NewServer(app, *cfg, logger, db, mongo)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err = s.Run(ctx); err != nil {
			logger.Fatalf("Run: %s", err)
		}
	}()

	<-ctx.Done()

	if err := app.Shutdown(); err != nil {
		logger.Fatalf("Shutdown: %s", err)
	}

	wg.Wait()
}
