package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	t.Run("the sum function returns sum of two strings", func(t *testing.T) {
		result, err := Sum("1", "1")
		assert.NoError(t, err)
		assert.Equal(t, result.Value, 2.0)
	})
	t.Run("the sum function returns an error if the first argument is invalid", func(t *testing.T) {
		result, err := Sum("a", "1")
		assert.Error(t, err)
		assert.Nil(t, result)
	})
	t.Run("the sum function returns an error if the second argument is invalid", func(t *testing.T) {
		result, err := Sum("1", "a")
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}
