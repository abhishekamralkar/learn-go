package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("without name and language", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertMessage(t, got, want)
	})

	t.Run("with name", func(t *testing.T) {
		got := Hello("Anay", "")
		want := "Hello, Anay"

		assertMessage(t, got, want)
	})

	t.Run("with language", func(t *testing.T) {
		got := Hello("", "Hindi")
		want := "Namaste, World"

		assertMessage(t, got, want)
	})

	t.Run("hindi", func(t *testing.T) {
		got := Hello("Anay", "Hindi")
		want := "Namaste, Anay"

		assertMessage(t, got, want)
	})

	t.Run("french", func(t *testing.T) {
		got := Hello("Anay", "French")
		want := "Bonjour, Anay"

		assertMessage(t, got, want)
	})

	t.Run("spanish", func(t *testing.T) {
		got := Hello("Anay", "Spanish")
		want := "Hola, Anay"

		assertMessage(t, got, want)
	})
}

func assertMessage(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
