package main

import "testing"

func TestSum(t *testing.T) {
	test := []struct {
		a, b, expected int
	}{
		{2, 2, 4},
		{3, 2, 5},
		{4, 2, 6},
		{5, 2, 7},
	}

	for _, test := range test {
		result := sum(test.a, test.b)

		if result != test.expected {
			t.Errorf("The result must be %d", test.expected)
		}
	}
}
