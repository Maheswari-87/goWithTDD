package main

import "fmt"

const french = "French"
const spanish = "Spanish"
const greetEnglish = "Hello, "
const greetSpanish = "Hola, "
const greetFrench = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}
	return greetPrefix(language) + name
}
func greetPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = greetFrench
	case spanish:
		prefix = greetSpanish
	default:
		prefix = greetEnglish
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
