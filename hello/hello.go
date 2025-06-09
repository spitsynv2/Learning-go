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
		if name == "" {
			name = "World"
		}
	case "esp":
		prefix = spanishHelloPrefix
		if name == "" {
			name = "Mundo"
		}
	default:
		prefix = englishHelloPrefix
		if name == "" {
			name = "World"
		}
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("world", "eng"))
}
