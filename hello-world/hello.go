package main

import "fmt"

const hindi = "Hindi"
const french = "French"
const spanish = "Spanish"
const englishPrefix = "Hello, "
const hindiPrefix = "Namaste, "
const spanishPrefix = "Hola, "
const frenchPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	switch language {
	case hindi:
		return hindiPrefix + name
	case french:
		return frenchPrefix + name
	case spanish:
		return spanishPrefix + name
	default:
		return englishPrefix + name

	}

}

func main() {
	fmt.Println(Hello("Anay", ""))
}
