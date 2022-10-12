package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBaseCut(t *testing.T) {
	given := "first test\tsecond test\tword\tquestion"
	expected := "first test"

	actual := cut(given, &flags{
		fields:    "0",
		delimiter: "\t",
	})

	assert.Equal(t, expected, actual)
}

func TestMultiplyFieldsCut(t *testing.T) {
	given := "first test\tsecond test\tword\tquestion"
	expected := "second testword"

	actual := cut(given, &flags{
		fields:    "1,2",
		delimiter: "\t",
	})

	assert.Equal(t, expected, actual)
}

func TestDifferentDelimiterCut(t *testing.T) {
	given := "first test\tsecond test\tword\tquestion"
	expected := "test\tsecondtest\tword\tquestion"

	actual := cut(given, &flags{
		fields:    "1,2",
		delimiter: " ",
	})

	assert.Equal(t, expected, actual)
}

func TestSeparatedCut(t *testing.T) {
	given := "first test\tsecond test\tword\tquestion"
	expected := ""

	actual := cut(given, &flags{
		fields:      "1,2",
		delimiter:   "\n",
		isSeparated: true,
	})

	assert.Equal(t, expected, actual)
}
