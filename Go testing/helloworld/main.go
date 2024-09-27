package main

import "fmt"

const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "

func Hello(name string, lang string) string {
	prefix := ""
	suffix := ""
	if lang == "English" || lang == "" {
		prefix = englishPrefix
		suffix = "World"
	} else {
		prefix = spanishPrefix
		suffix = "Mundo"
	}
	if name == "" {
		return prefix + suffix
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("Federico", ""))
}
