package main

import "log"

func main() {
	AssertEqual("CJ", "CJ")
	AssertEqual(1, "1")
}

func AssertEqual(got, want interface{}) {
	if got != want {
		log.Fatalf("FAILED: got %+v, want %+v", got, want)
	} else {
		log.Printf("PASSED: %+v did equal %+v\n", got, want)
	}
}

func AssertNotEqual(got, want interface{}) {
	if got == want {
		log.Fatalf("FAILED: got %+v, want %+v", got, want)
	} else {
		log.Printf("PASSED: %+v did not equal %+v\n", got, want)
	}
}
