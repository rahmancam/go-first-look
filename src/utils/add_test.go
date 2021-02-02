package utils

import "testing"

func TestAdd(t *testing.T) {
	expected := 9
	actual := Add(5, 4)
	if actual != expected {
		t.Errorf("Add function does not add up: Expected: %d, Actual: %d", expected, actual)
	}
}
