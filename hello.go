package main

import "fmt"

const englishPrefix = "Hello, %s"
const spanishPrefix = "Hola, %s"
const frenchPrefix = "Bonjour, %s"

func greetingPrefix(language string) string {
	switch language {
	case "Spanish":
		return spanishPrefix
	case "French":
		return frenchPrefix
	default:
		return englishPrefix
	}
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return fmt.Sprintf(greetingPrefix(language), name)
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
