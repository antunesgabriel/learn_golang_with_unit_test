package ints

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	result := Sum(2, 2)
	expected := 4

	if (result != expected) {
		t.Errorf("Expected: '%d' --- Result: '%d'", expected, result)
	}
}

func ExampleSum() {
	result := Sum(34, 26)

	fmt.Println(result)
	// Output: 60
}
