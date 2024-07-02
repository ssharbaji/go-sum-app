package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	assert := assert.New(t)

	result := Sum(3, 4)
	assert.Equal(7, result, "Should be equal")
}
