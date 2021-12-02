package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	result, err := Sum("1", "1")
	assert.NoError(t, err)
	assert.Equal(t, result.Value, 2.0)
}
