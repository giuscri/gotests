package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "english")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, world' when passed an empty string", func(t *testing.T) {
		got := Hello("", "english")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to people in spanish", func(t *testing.T) {
		got := Hello("Chris", "spanish")
		want := "Hola, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to people in french", func(t *testing.T) {
		got := Hello("Chris", "french")
		want := "Bonjour, Chris"
		assertCorrectMessage(t, got, want)
	})
}
