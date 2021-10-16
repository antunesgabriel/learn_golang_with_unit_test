package iteration

import (
	"fmt"
	"testing"
)

func checkResult(t *testing.T, expected *string, result *string) {
	if *expected != *result {
		t.Errorf("Expected '%s' --- Result: '%s'", *expected, *result)
	}
}

func TestRepeat(t *testing.T) {
	t.Run("It should repeat word 4 times", func(t *testing.T) {
		result := Repeat("a", 4)
		expected := "aaaa"

		checkResult(t, &expected, &result)
	})

	t.Run("It should repeat word 6 times", func(t *testing.T) {
		result := Repeat("vasco", 6)
		expected := "vascovascovascovascovascovasco"

		checkResult(t, &expected, &result)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("g", 7)
	}
}

func ExampleRepeat() {
	result := Repeat("antunes", 5)

	fmt.Println(result)
	// Output: antunesantunesantunesantunesantunes
}
