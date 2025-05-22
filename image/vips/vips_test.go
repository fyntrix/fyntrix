package vips

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVipsVersion(t *testing.T) {
	vip := New()
	require.NotNil(t, vip)

	ver := vip.Version()
	assert.NotEmpty(t, ver)
	t.Logf("Vips version: %s", ver)
}
