package main

import "testing"

func Test(t *testing.T) {

	t.Run("Hello to a person", func(t *testing.T) {
		got := Hello("Federico", "English")
		want := "Hello, Federico"

		assertCorretMessage(t, want, got)
	})

	t.Run("Hello to anyone", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		assertCorretMessage(t, want, got)
	})

	t.Run("Hello in spanish to anyone", func(t *testing.T) {
		got := Hello("", "Spanish")
		want := "Hola, Mundo"

		assertCorretMessage(t, want, got)
	})

	t.Run("Hello in spanish to person", func(t *testing.T) {
		got := Hello("Marika", "Spanish")
		want := "Hola, Marika"

		assertCorretMessage(t, want, got)
	})
	t.Run("Hello in spanish to person", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorretMessage(t, want, got)
	})
	t.Run("Hello in spanish to person", func(t *testing.T) {
		got := Hello("Federico", "")
		want := "Hello, Federico"

		assertCorretMessage(t, want, got)
	})
}

func assertCorretMessage(t *testing.T, want string, got string) {
	t.Helper()
	if got != want {
		t.Errorf("Want %q, got %q", want, got)
	}
}
