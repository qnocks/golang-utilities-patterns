package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnagram(t *testing.T) {
	given := []string{"Тяпка", "пятак", "лИсток", "столик", "пятка", "слиток"}
	expected := map[string][]string{
		"листок": {"листок", "слиток", "столик"},
		"пятак":  {"пятак", "пятка", "тяпка"},
	}

	actual := search(&given)
	assert.Equal(t, &expected, actual)
}
