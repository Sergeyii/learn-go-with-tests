package main

import "fmt"

func HelloWorld() string {
	return "Hello world!"
}

func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("World"))
}
