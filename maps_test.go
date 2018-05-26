package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlipMapUint64(t *testing.T) {
	original := make(map[uint64]uint64)
	original[1] = 2
	original[3] = 4
	original[5] = 6

	flipped := FlipMapUint64(original)

	assert.True(t, flipped[2] == 1)
	assert.True(t, flipped[4] == 3)
	assert.True(t, flipped[6] == 5)
}
