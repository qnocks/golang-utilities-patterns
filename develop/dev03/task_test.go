package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumSort(t *testing.T) {
	given := []string{"22222", "1111", "11111", "3333", "55555"}
	expected := []string{"1111", "11111", "22222", "3333", "55555"}

	actual := processSort(given, &flags{column: -1, hasNumber: true})
	assert.Equal(t, expected, actual)
}

func TestBaseSort(t *testing.T) {
	given := []string{"aaaaa", "aaaaa", "sdadas", "posda", "nhngn"}
	expected := []string{"aaaaa", "nhngn", "posda", "sdadas"}

	actual := processSort(given, &flags{column: -1, isUnique: true})
	assert.Equal(t, expected, actual)
}

func TestColumnSort(t *testing.T) {
	given := []string{"aaaaa bbbbb", "aaaaa cccccc", "sdadas ojsa", "posda sdadsad", "nhngn zzzzz"}
	expected := []string{"nhngn zzzzz", "posda sdadsad", "sdadas ojsa", "aaaaa cccccc", "aaaaa bbbbb"}

	actual := processSort(given, &flags{column: 1, isReverse: true})
	assert.Equal(t, expected, actual)
}
