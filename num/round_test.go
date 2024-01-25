package num

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRound(t *testing.T) {
	assert.Equal(t, Round(230.55), 231, "230.55 = 231")
	assert.Equal(t, RoundUp(230.55), 231, "230.55 = 231")
	assert.Equal(t, RoundDown(230.55), 230, "230.55 = 230")
}
