package config

import (
	"strings"

	"github.com/fyntrix/fyntrix/image"
	"github.com/fyntrix/fyntrix/pkg/logger"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	Development bool           `koanf:"development"`
	Image       *image.Config  `koanf:"image"`
	Logger      *logger.Config `koanf:"logger"`
}

func Load(path string) (*Config, error) {
	cfg := _default()

	koan := koanf.New(".")

	// Load from file
	if err := koan.Load(file.Provider(path), toml.Parser()); err != nil {
		return nil, err
	}

	// Load from env (after file)
	if err := koan.Load(env.Provider("FYNTRIX_", ".", func(key string) string {
		key = strings.TrimPrefix(key, "FYNTRIX_")
		key = strings.ToLower(key)
		parts := strings.Split(key, "_")
		if len(parts) > 1 {
			return parts[0] + "." + strings.Join(parts[1:], "_")
		}

		return key
	}), nil); err != nil {
		return nil, err
	}

	// ✅ Now apply everything to the struct
	if err := koan.Unmarshal("", cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) BasicCheck() error {
	if err := c.Image.BasicCheck(); err != nil {
		return err
	}

	return c.Logger.BasicCheck()
}

func _default() *Config {
	return &Config{
		Image:  image.DefaultConfig(),
		Logger: logger.DefaultConfig(),
	}
}
