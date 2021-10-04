package main

import "testing"

func TestHello(t *testing.T) {
	checkResult := func (t *testing.T, expected string, result string) {
		t.Helper()

		if result != expected {
			t.Errorf("Expected: %s - But received: %s", expected, result)
		}
	}

	t.Run("say hello to people", func (t *testing.T) {
		name := "Gabriel"
		expected := "Hello, " + name
		result := Hello(name, "")

		checkResult(t, expected, result)
	})

	t.Run("Say hello word when the name is empty", func (t *testing.T) {
		name := ""
		expected := "Hello, Word"
		result := Hello(name, "")

		checkResult(t, expected, result)
	})

	t.Run("Say hello in spanish when you receive spanish language", func (t *testing.T) {
		language := "spanish"
		name := "Pablo"
		expected := "Hola, " + name

		result := Hello(name, language)

		checkResult(t, expected, result)
	})

	t.Run("Say hello in french when you get french", func (t *testing.T) {
		language := "french"
		name := "Grizlley"
		expected := "Bonjour, " + name

		result := Hello(name, language)

		checkResult(t, expected, result)
	})

	t.Run("Say hello in italian when you get italian", func (t *testing.T) {
		language := "italian"
		name := "Ramon"
		expected := "Ciao, " + name

		result := Hello(name, language)

		checkResult(t, expected, result)
	})
}
