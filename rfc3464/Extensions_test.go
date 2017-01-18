package rfc3464

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtensions(t *testing.T) {
	e := make(Extensions)
	assert.Len(t, e, 0)

	e.Set("test-key", "value")
	assert.Len(t, e, 1)
	assert.Contains(t, e, "Test-Key")
	assert.Equal(t, "value", e["Test-Key"])
	assert.Equal(t, "value", e.Get("TEST-KEY"))

	e.Del("Test-Key")
	assert.Len(t, e, 0)
	assert.Equal(t, "", e.Get("test-key"))
}

func TestExtensionsNil(t *testing.T) {
	var e Extensions

	assert.Equal(t, "", e.Get("woot"))
}
