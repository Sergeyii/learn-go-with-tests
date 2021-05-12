package main

import "fmt"

const englishHelloPrefix = "Hello, "

func HelloWorld() string {
	return "Hello world!"
}

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("World"))
}
