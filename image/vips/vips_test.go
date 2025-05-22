package vips

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVipsInitAndClose(t *testing.T) {
	v := New(&Config{
		ConcurrentLevel: 1,
		MaxCacheSize:    100,
	})

	require.False(t, v.Inited())

	err := v.Init()
	require.NoError(t, err)
	assert.True(t, v.Inited(), "vips should be initialized")

	err = v.Close()
	require.NoError(t, err)
	assert.False(t, v.Inited(), "vips should be shut down")
}

func TestVipsVersion(t *testing.T) {
	vip := New(DefaultConfig())
	require.NotNil(t, vip)

	ver := vip.Version()
	assert.NotEmpty(t, ver)
	t.Logf("Vips version: %s", ver)
}
