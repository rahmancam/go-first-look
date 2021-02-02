package utils

import "testing"

func TestMakeExcited(t *testing.T) {
	expected := "OMG SO EXCITING!"
	actual := MakeExcited("omg so exciting")
	if actual != expected {
		t.Errorf("MakeExcited was incorrect! Expected : %s, Actual: %s", expected, actual)
	}
}
