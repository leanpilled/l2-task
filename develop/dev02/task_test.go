package main

import "testing"

func TestUnpackString(t *testing.T) {
	testTable := []struct {
		str      string
		expected string
	}{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"qwe\\4\\5", "qwe45"},
		{"qwe\\45", "qwe44444"},
		{"qwe\\\\5", "qwe\\\\\\\\\\"},
	}

	for _, test := range testTable {
		result := unpackString(test.str)
		if result != test.expected {
			t.Errorf("Expected: %s, got: %s", test.expected, result)
		}
	}

}
