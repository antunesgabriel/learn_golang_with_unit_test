package main

import "fmt"

const helloDefault = "Hello"
const helloSpanish = "Hola"
const helloFrench = "Bonjour"
const helloItalian = "Ciao"

const spanish = "spanish"
const french = "french"
const italian = "italian"

func Hello(name string, language string) string {
	if name == "" {
		name = "Word"
	}

	return switchPrefix(language) +  ", " + name
}

func switchPrefix (language string) (prefix string) {
	switch language {
	case spanish:
		prefix = helloSpanish
	case french:
		prefix = helloFrench
	case italian:
		prefix = helloItalian
	default:
		prefix = helloDefault
	}
	return
}

func main() {
	fmt.Println(Hello("Word", ""))
}
