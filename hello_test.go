package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("in Kazakh", func(t *testing.T) {
		got := Hello("Касым", "Kazakh")
		want := "Ассаламу алейкум Касым"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Епта", "French")
		want := "Bonjour Епта"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Go", "")
		want := "Hello Go"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello world"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
