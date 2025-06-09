package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("hello to people", func(t *testing.T) {
		got := Hello("people")
		expected := "Hello, people"

		assertCorrectMessage(t, got, expected)

	})
	t.Run("hello to world", func(t *testing.T) {
		got := Hello("world")
		expected := "Hello, world"

		assertCorrectMessage(t, got, expected)

	})
	t.Run("with empty string", func(t *testing.T) {
		got := Hello("world")
		expected := "Hello, world"
		assertCorrectMessage(t, got, expected)
	})
}

func assertCorrectMessage(t testing.TB, got, expected string) {
	t.Helper()
	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}
}
