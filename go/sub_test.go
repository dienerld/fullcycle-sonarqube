package main

import "testing"

func TestSub(t *testing.T) {
	// make array tests

	tests := []struct {
		a, b, expected int
	}{
		{2, 2, 0},
		{3, 2, 1},
		{4, 2, 2},
		{5, 2, 3},
	}

	for _, test := range tests {
		result := sub(test.a, test.b)

		if result != test.expected {
			t.Errorf("The result must be %d", test.expected)
		}
	}
}
