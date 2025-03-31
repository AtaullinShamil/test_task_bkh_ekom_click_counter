package main

import (
	"context"
	"os"

	"github.com/AtaullinShamil/test_task_bkh_ekom_click_counter/config"
	"github.com/AtaullinShamil/test_task_bkh_ekom_click_counter/pkg/postgres"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()

	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	cfg, err := config.Load()
	if err != nil {
		logger.Fatalf("Load: %s", err)
	}

	migrateDown := false

	for _, v := range os.Args {
		if v == "--down" {
			migrateDown = true
		}
	}

	migrate(ctx, logger, &cfg.Postgres, migrateDown)
}

func migrate(ctx context.Context, logger *logrus.Logger, postgresConfig *postgres.Config, migrateDown bool) {
	db, err := postgres.NewPostgresClient(ctx, *postgresConfig)
	if err != nil {
		logger.Fatalf("NewPostgresClient: %s", err)
	}

	defer db.Close()

	if err = goose.SetDialect("postgres"); err != nil {
		logger.Fatalf("SetDialect: %s", err)
	}

	if migrateDown {
		if err = goose.Down(db.DB, "./migrations"); err != nil {
			logger.Fatalf("Down: %s", err)
		}

		return
	}

	if err = goose.Up(db.DB, "./migrations"); err != nil {
		logger.Fatalf("Up: %s", err)
	}
}
