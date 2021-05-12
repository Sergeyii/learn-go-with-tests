package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got=%q, want=%q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		//Получил по факту
		got := Hello("Chris", "")
		//Ожидаю
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		//Получил по факту
		got := Hello("", "")
		//Ожидалось
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		//Получено
		got := Hello("Elodie", "Spanish")
		//Ожидалось
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		//Получено
		got := Hello("Marta", "French")
		//Ожидалось
		want := "Bonjour, Marta"

		assertCorrectMessage(t, got, want)
	})
}
