package main

import "fmt"

const hello = "Hello"

func Hello(name string) string {
	return hello + ", " + name
}

func main() {
	fmt.Println(Hello("Word"))
}
