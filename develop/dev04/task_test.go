package main

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	testTable := []struct {
		dict     []string
		expected map[string][]string
	}{
		{
			[]string{"пятак", "ПЯТКА", "тяпка", "столик", "слиток", "листок"},
			map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"столик": {"листок", "слиток", "столик"},
			},
		},
		{
			[]string{"пятак"},
			map[string][]string{},
		},
	}

	for _, test := range testTable {
		result := findAnagrams(test.dict)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Expected: %s, got: %s", test.expected, result)
		}
	}
}
