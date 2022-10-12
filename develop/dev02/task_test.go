package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_unbox(t *testing.T) {
	given := []string{"a4bc2d5e", "abcd", "45", ""}
	expected := []string{"aaaabccddddde", "abcd", "", ""}
	actual := make([]string, len(expected))

	for i := 0; i < len(expected); i++ {
		actual[i] = unbox(given[i])
	}

	assert.Equal(t, expected, actual)
}
