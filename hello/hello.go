package main

import "fmt"

const (
	defaultName        = "World"
	engHelloPrefix     = "Hello, "
	french             = "French"
	frenchHelloPrefix  = "Bonjour, "
	spanish            = "Spanish"
	spanishHelloPrefix = "Hola, "
)

func Hello(name string, language string) string {
	return greetingPrefix(language) + greetingName(name)
}

func greetingName(name string) (greetingName string) {
	if name == "" {
		return defaultName
	}

	return name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		return frenchHelloPrefix
	case spanish:
		return spanishHelloPrefix
	default:
		return engHelloPrefix
	}
}

func main() {
	fmt.Println(Hello("world", ""))
}
