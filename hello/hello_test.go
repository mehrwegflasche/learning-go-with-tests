package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("expected %d, but got %d", want, got)
		}
	}

	t.Run("Say Hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say Hello, World when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello in Spanish", func(t *testing.T) {
		got := Hello("Chester", "Spanish")
		want := "Hola, Chester"
		assertCorrectMessage(t, got, want)
	})
}
