package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("hello to people, empty language", func(t *testing.T) {
		got := Hello("people", "")
		expected := "Hello, people"

		assertCorrectMessage(t, got, expected)

	})
	t.Run("hello to people, eng", func(t *testing.T) {
		got := Hello("people", "eng")
		expected := "Hello, people"

		assertCorrectMessage(t, got, expected)
	})
	t.Run("hello to world", func(t *testing.T) {
		got := Hello("World", "")
		expected := "Hello, World"

		assertCorrectMessage(t, got, expected)

	})
	t.Run("with empty string", func(t *testing.T) {
		got := Hello("", "")
		expected := "Hello, World"
		assertCorrectMessage(t, got, expected)
	})
	t.Run("with empty string, eng", func(t *testing.T) {
		got := Hello("", "eng")
		expected := "Hello, World"
		assertCorrectMessage(t, got, expected)
	})
	t.Run("with empty string, spanish", func(t *testing.T) {
		got := Hello("", "esp")
		expected := "Hola, Mundo"
		assertCorrectMessage(t, got, expected)
	})
}

func assertCorrectMessage(t testing.TB, got, expected string) {
	t.Helper()
	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}
}
