package main

import "testing"

func TestSub(t *testing.T) {
	result := sub(2, 2)

	if result != 0 {
		t.Error("The result must be 0")
	}
}
