package config

import (
	"github.com/AtaullinShamil/test_task_bkh_ekom_click_counter/pkg/mongo"
	"github.com/AtaullinShamil/test_task_bkh_ekom_click_counter/pkg/postgres"
)

type Config struct {
	Server   Server          `mapstructure:"server"`
	Postgres postgres.Config `mapstructure:"postgres"`
	Mongo    mongo.Config    `mapstructure:"mongo"`
}

type Server struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
