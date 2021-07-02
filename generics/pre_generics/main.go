package main

import "log"

type StackOfInts = Stack
type StackOfStrings = Stack

type Stack struct {
	values []interface{}
}

func (s *Stack) Push(value interface{}) {
	s.values = append(s.values, value)
}

func (s *Stack) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		var zero interface{}
		return zero, false
	}

	index := len(s.values) - 1
	el := s.values[index]
	s.values = s.values[:index]
	return el, true
}

func main(){
	//INT stack
	myStackOfInts := new(StackOfInts)

	myStackOfInts.Push(1)
	myStackOfInts.Push(2)
	firstNum, _ := myStackOfInts.Pop()
	secondNum, _ := myStackOfInts.Pop()

	//get our ints from out interface{}
	reallyFirstNum, ok := firstNum.(int)
	AssertTrue(ok) //need to check we definitely got an int out of the interface{}

	reallySecondNum, ok := secondNum.(int)
	AssertTrue(ok) //and again!

	AssertEqual(reallyFirstNum+reallySecondNum, 3)

	//STRING stack
	//myStackOfStrings := new(StackOfStrings)
}

func AssertTrue(thing bool) {
	if thing {
		log.Printf("PASSED: Expected thing to be true and it was\n")
	}else{
		log.Fatalf("FAILED: expected true but got false")
	}
}

func AssertFalse(thing bool) {
	if !thing {
		log.Printf("PASSED: Expected thing to be false and it was\n")
	}else{
		log.Fatalf("FAILED: expected false but got true")
	}
}

func AssertEqual[T comparable](got, want T) {
	if got != want {
		log.Fatalf("FAILED: got %+v, want %+v", got, want)
	}else{
		log.Printf("PASSED: %+v did equal %+v\n", got, want)
	}
}

func AssertNotEqual[T comparable](got, want T) {
	if got == want {
		log.Fatalf("FAILED: got %+v, want %+v", got, want)
	}else{
		log.Printf("PASSED: %+v did not equal %+v\n", got, want)
	}
}