package image

import "github.com/fyntrix/fyntrix/image/vips"

type Config struct {
	Vips *vips.Config `koanf:"vips"`
}

func DefaultConfig() *Config {
	return &Config{
		Vips: vips.DefaultConfig(),
	}
}

func (c *Config) BasicCheck() error {
	return c.Vips.BasicCheck()
}
