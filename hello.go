package main

import "fmt"

const spanish = "spanish"
const french = "french"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func Hello(name, language string) string {
	prefix := greetingPrefix(language)

	if len(name) == 0 {
		return prefix + "world"
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", "english"))
}
