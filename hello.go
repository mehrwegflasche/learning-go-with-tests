package main

import "fmt"

const englishPrefix = "Hello, %s"

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf(englishPrefix, name)
}

func main() {
	fmt.Println(Hello("Chris"))
}
