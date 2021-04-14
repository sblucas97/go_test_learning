package main

import "fmt"

// English const to use at package level
const English string = "English"
// Portuguese const to use at package level
const Portuguese string = "Portuguese"
// Spanish const to use at package level
const Spanish string = "Spanish"
// EnglishGreeting const to use at package level
const EnglishGreeting = "Hello, "
// PortugueseGreeting const to use at package level
const PortugueseGreeting = "Olá, "
// SpanishGreeting const to use at package level
const SpanishGreeting = "Hola, "

// Hello returns a greeting
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingMessage(language) + name
}

func greetingMessage(language string) string{
	switch language {
	case Portuguese:
		return PortugueseGreeting
	case Spanish:
		return SpanishGreeting
	default:
		return EnglishGreeting
	}
}

func main() {
	fmt.Println(Hello("Lucas", ""))
}