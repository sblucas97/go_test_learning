package main

import "testing"

const name string = "Lucas"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("sayig hello to people", func(t *testing.T) {
		got := Hello(name, "")
		want := EnglishGreeting + name

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := EnglishGreeting + "World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Portuguese", func(t *testing.T) {
		got := Hello(name, Portuguese)
		want := PortugueseGreeting + name

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello(name, Spanish)
		want := SpanishGreeting + name

		assertCorrectMessage(t, got, want)
	})
}
