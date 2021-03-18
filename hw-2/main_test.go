package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFrequencyAnalysis(t *testing.T) {
	type pair struct {
		input  string
		output []string
	}
	test := []pair{
		{"Hello World Hello Hello World Gunt Gunt Gunt Gunt Gunt Hello", []string{"Gunt", "Hello", "World"}},
	}
	for _, pair := range test {
		require.Equal(t, FrequencyAnalysis(pair.input), pair.output, "Error with Unpack function")
	}
}
