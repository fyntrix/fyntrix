package logger

type Config struct {
	Colorful           bool              `koanf:"colorful"`
	MaxBackups         int               `koanf:"max_backups"`
	RotateLogAfterDays int               `koanf:"rotate_log_after_days"`
	Compress           bool              `koanf:"compress"`
	Targets            []string          `koanf:"targets"`
	Levels             map[string]string `koanf:"levels"`
}

func DefaultConfig() *Config {
	conf := &Config{
		Levels:             make(map[string]string),
		Colorful:           true,
		MaxBackups:         0,
		RotateLogAfterDays: 1,
		Compress:           true,
		Targets:            []string{"console"},
	}

	conf.Levels["default"] = "info"
	conf.Levels["_rest"] = "info"
	conf.Levels["_image"] = "warn"
	conf.Levels["_video"] = "warn"

	return conf
}

func (*Config) BasicCheck() error {
	return nil
}
