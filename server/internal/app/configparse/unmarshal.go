package configparse

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Options struct {
	DisableEnvironment bool
	From               struct {
		FilePath string
		Map      map[string]any
	}
}

func FromMap(m map[string]any) Opt {
	return func(o *Options) {
		o.From.Map = m
	}
}

func FromPath(path string) Opt {
	return func(o *Options) {
		o.From.FilePath = path
	}
}

type Opt = func(*Options)

func Unmarshal(v any, opts ...Opt) error {
	var opt Options
	for _, o := range opts {
		o(&opt)
	}
	cfg := viper.New()
	if opt.From.Map != nil {
		for key, value := range opt.From.Map {
			cfg.SetDefault(key, value)
		}
	}

	if !opt.DisableEnvironment {
		cfg.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		cfg.AutomaticEnv()
	}

	var err error
	if opt.From.FilePath != "" {
		cfg.AddConfigPath(".")
		cfg.SetConfigType("yaml")
		cfg.SetConfigName(opt.From.FilePath)
		err = cfg.ReadInConfig()
		if err != nil {
			return fmt.Errorf("reading config file: %w", err)
		}
	}

	err = cfg.Unmarshal(v)
	if err != nil {
		return fmt.Errorf("unmarshal config: %w", err)
	}

	return nil
}
