package zapdriver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceContext(t *testing.T) {
	t.Parallel()

	got := ServiceContext("test service name", "v1.0.0").Interface.(*serviceContext)

	assert.Equal(t, "test service name", got.Name)
	assert.Equal(t, "v1.0.0", got.Version)
}

func TestNewServiceContext(t *testing.T) {
	t.Parallel()

	got := newServiceContext("test service name", "v1.0.0")

	assert.Equal(t, "test service name", got.Name)
	assert.Equal(t, "v1.0.0", got.Version)
}
