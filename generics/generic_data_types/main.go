package main

import "log"

type StackOfInts struct {
	values []int
}

func (s *StackOfInts) Push(value int) {
	s.values = append(s.values, value)
}

func (s *StackOfInts) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *StackOfInts) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	index := len(s.values) - 1
	el := s.values[index]
	s.values = s.values[:index]
	return el, true
}

type StackOfStrings struct {
	values []string
}

func (s *StackOfStrings) Push(value string) {
	s.values = append(s.values, value)
}

func (s *StackOfStrings) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *StackOfStrings) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	index := len(s.values) - 1
	el := s.values[index]
	s.values = s.values[:index]
	return el, true
}

func main(){
	//INT stack
	myStackOfInts := new(StackOfInts)

	//check stack is empty
	AssertTrue(myStackOfInts.IsEmpty())

	//Add a thing, then check it's not empty
	myStackOfInts.Push(123)
	AssertFalse(myStackOfInts.IsEmpty())

	//add another thing, pop it back again
	myStackOfInts.Push(456)
	value, _ := myStackOfInts.Pop()
	AssertEqual(value, 456)
	value, _ = myStackOfInts.Pop()
	AssertEqual(value, 123)
	AssertTrue(myStackOfInts.IsEmpty())


	//STRING stack
	myStackOfStrings := new(StackOfStrings)

	//check stack is empty
	AssertTrue(myStackOfStrings.IsEmpty())

	//add a thing, then check it's not empty
	myStackOfStrings.Push("one two three")
	AssertFalse(myStackOfStrings.IsEmpty())

	//add another thing, pop it back again
	myStackOfStrings.Push("four five six")
	strValue, _ := myStackOfStrings.Pop()
	AssertEqual(strValue, "four five six")
	strValue, _ = myStackOfStrings.Pop()
	AssertEqual(strValue, "one two three")
	AssertTrue(myStackOfStrings.IsEmpty())
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