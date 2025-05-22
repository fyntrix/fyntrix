package vips

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVipsVersion(t *testing.T) {
	vip := New()
	ver := vip.Version()

	require.NotNil(t, vip)
	assert.NotEmpty(t, ver)
	t.Logf("Vips version: %s", ver)
}
