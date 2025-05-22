package vips

type Config struct {
	ConcurrentLevel int  `koanf:"concurrent_level"`
	MaxCacheMem     int  `koanf:"max_cache_mem"`
	MaxCacheSize    int  `koanf:"max_cache_size"`
	MaxCacheFiles   int  `koanf:"max_cache_files"`
	ReportLeaks     bool `koanf:"report_leaks"`
	CacheTrace      bool `koanf:"cache_trace"`
}

func DefaultConfig() *Config {
	return &Config{
		ConcurrentLevel: 1,
		MaxCacheMem:     50 * 1024 * 1024,
		MaxCacheSize:    100,
		MaxCacheFiles:   0,
		ReportLeaks:     true,
		CacheTrace:      true,
	}
}

func (c *Config) BasicCheck() error {
	return nil
}
