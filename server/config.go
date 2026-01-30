package server

import "github.com/philoj/goplanes/server/internal/app/api"

type Config struct {
	AutoMigrate bool `mapstructure:"auto_migrate"`

	Server api.Config `mapstructure:"server"`
}
