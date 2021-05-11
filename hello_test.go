package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		//Получил по факту
		got := Hello("Chris")
		//Ожидаю
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got=%q, want=%q", got, want)
		}
	})

	t.Run("saying 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		//Получил по факту
		got := Hello("World")
		//Ожидалось
		want := "Hello, World"

		if got != want {
			t.Errorf("got=%q, want=%q", got, want)
		}
	})
}
