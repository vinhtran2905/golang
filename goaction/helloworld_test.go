package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	tt := []struct {
		input    string
		expected string
	}{
		{"", "Hello, world"},
		{"Michael", "Hello, Michael"},
	}

	for _, tc := range tt {
		actual := Hello(tc.input)
		if actual != tc.expected {
			t.Errorf("Failed. Expected %s, got %s", tc.expected, actual)
		}
	}
}
