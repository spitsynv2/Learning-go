package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, lang string) string {
	prefix := ""

	switch lang {
	case "eng":
		prefix = englishHelloPrefix
		checkName(&name, "World")
	case "esp":
		prefix = spanishHelloPrefix
		checkName(&name, "Mundo")
	default:
		prefix = englishHelloPrefix
		checkName(&name, "World")
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("world", "eng"))
}

func checkName(name *string, defaultValue string) {
	if *name == "" {
		*name = defaultValue
	}
}
