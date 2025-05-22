package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	cfg, err := Load("./config.example.toml")
	require.NoError(t, err)
	require.NoError(t, cfg.BasicCheck())

	require.NotNil(t, cfg)
	require.NotNil(t, cfg.Image)
}
