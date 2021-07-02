package main

import "log"

type StackOfInts = Stack
type StackOfStrings = Stack

type Stack[T any] struct {
	values []T
}

func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	index := len(s.values) - 1
	el := s.values[index]
	s.values = s.values[:index]
	return el, true
}

func main(){
	myStackOfInts := new(Stack[int])

	//check stack is empty
	AssertTrue(myStackOfInts.IsEmpty())

	//add a thing, then check it's not empty
	myStackOfInts.Push(123)
	AssertFalse(myStackOfInts.IsEmpty())

	//add another thing, pop it back again
	myStackOfInts.Push(456)
	value, _ := myStackOfInts.Pop()
	AssertEqual(value, 456)
	value, _ = myStackOfInts.Pop()
	AssertEqual(value, 123)
	AssertTrue(myStackOfInts.IsEmpty())

	//can get the numbers we put in as number, not untyped interface{}
	myStackOfInts.Push(1)
	myStackOfInts.Push(2)
	firstNum, _ := myStackOfInts.Pop()
	secondNum, _ := myStackOfInts.Pop()
	AssertEqual(firstNum+secondNum, 3)
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