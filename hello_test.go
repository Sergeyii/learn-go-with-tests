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
