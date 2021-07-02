package main

import "log"

func main(){
	AssertEqual(1, 1)
	AssertEqual("1", "1")
	AssertEqual(1, 2)
	//AssertEqual(1, "1") //- uncomment me to see compilation error
}

func AssertEqual[T any](got, want T){
	if got != want {
		log.Fatalf("FAILED: got %+v, want %+v", got, want)
	}else{
		log.Printf("PASSED: %+v did equal %+v\n", got, want)
	}
}

func AssertNotEqual[T any](got, want T){
	if got == want {
		log.Fatalf("FAILED: got %+v, want %+v", got, want)
	}else{
		log.Printf("PASSED: %+v did not equal %+v\n", got, want)
	}
}