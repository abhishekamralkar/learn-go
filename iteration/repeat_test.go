package main

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("pass both character and repeastcount ", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assertMessage(t, repeated, expected)
	})

	t.Run("Pass only character", func(t *testing.T) {
		repeated := Repeat("b", 0)
		expected := "b"

		assertMessage(t, repeated, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 1000)
	}
}

func assertMessage(t testing.TB, repeated, expected string) {
	t.Helper()

	if repeated != expected {
		t.Errorf("expexted %q but got %q", expected, repeated)
	}
}
