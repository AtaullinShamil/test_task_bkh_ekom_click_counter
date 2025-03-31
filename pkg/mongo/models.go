package mongo

import "time"

type Config struct {
	URI            string        `mapstructure:"uri"`
	Username       string        `mapstructure:"username"`
	Password       string        `mapstructure:"password"`
	ConnectTimeout time.Duration `mapstructure:"connectTimeout"`
}
