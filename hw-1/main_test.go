package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdditional(t *testing.T) {
	require.Equal(t, Addition(2, 2), 4, "2 + 2 to equal 4")
}
func TestUnpack(t *testing.T) {
	type pair struct {
		input  string
		output string
	}
	test := []pair{
		{"abcd", "abcd"},
		{"abc", "abc"},
		{"a4bc2d5e", "aaaabccddddde"},
		{"45", ""},
		{"a4", "aaaa"},
	}
	for _, pair := range test {
		require.Equal(t, Unpack(pair.input), pair.output, "Error with Unpack function")
	}
}
