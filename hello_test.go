package main

import "testing"

func TestHelloWorld(t *testing.T) {
	//Ожидаю
	expected := "Hello world!"

	//Получилось на самом деле
	got := HelloWorld()

	if expected != got {
		t.Errorf("Expected =%q, got=%q", expected, got)
	}
}

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}
}
