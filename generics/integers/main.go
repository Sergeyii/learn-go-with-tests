package main

import "log"

func main() {
	AssertEqual(1, 1)
	AssertNotEqual(1, 21)

	AssertEqual(50, 100) //this should fail
	AssertNotEqual(2, 2) //so you wont see this print
}

func AssertEqual(got, want int) {
	if got != want {
		log.Fatalf("FAILED: got %d, want %d", got, want)
	} else {
		log.Printf("PASSED: %d did equal %d\n", got, want)
	}
}

func AssertNotEqual(got, want int) {
	if got == want {
		log.Fatalf("FAILED: got %d, want %d", got, want)
	} else {
		log.Printf("PASSED: %d did not equal %d\n", got, want)
	}
}
