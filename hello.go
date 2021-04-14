package main

import "fmt"

// EnglishGreeting const to use at package level
const EnglishGreeting = "Hello, "

// Hello returns a greeting
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return EnglishGreeting + name
}

func main() {
	fmt.Println(Hello("Lucas"))
}