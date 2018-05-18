package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}

	index := Find(len(slice), func(i int) bool { return slice[i] == 3 })
	assert.Equal(t, index, 2)

	index = Find(len(slice), func(i int) bool { return slice[i] == 9 })
	assert.Equal(t, index, -1)
}
