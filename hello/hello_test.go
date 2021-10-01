package main

import "testing"

func TestHello(t *testing.T) {
	name := "Gabriel"
	expected := "Hello, " + name
	result := Hello(name)

	if result != expected {
		t.Errorf("Expected: %s - But received: %s", expected, result)
	}
}
