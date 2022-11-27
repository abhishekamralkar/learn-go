package main

import "testing"

func TestTotal(t *testing.T) {
	t.Run("pass both array and count", func(t *testing.T) {
		got := Total([5]int{34, 56, 78, 94, 100})
		want := 362

		assertMessage(t, got, want)
	})
}

func assertMessage(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
