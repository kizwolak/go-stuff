package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name, lang string) string {
	name = checkName(name)
	return checkLanguage(lang) + name
}

func main() {
	fmt.Println(Hello("World", ""))
}

func checkName(name string) string {
	if name == "" {
		return "World"
	}
	return name
}

func checkLanguage(lang string) string {
	switch lang {
	case "Spanish":
		return spanishHelloPrefix
	case "French":
		return "Bonjour, "
	case "Polish":
		return "Cześć, "
	default:
		return englishHelloPrefix
	}
}
