package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("find a known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("find an un-known word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		want := ErrNotFound
		assertError(t, got, want)
	})

	t.Run("add a word", func(t *testing.T) {
		dictionary.Add("another", "this is a test")
		want := "this is a test"
		got, err := dictionary.Search("another")
		if err != nil {
			t.Fatal("should find added word", err)
		}
		assertStrings(t, got, want)
	})

	t.Run("add an existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")
		assertError(t, err, ErrWordExists)
	})

	t.Run("update an existing word", func(t *testing.T) {
		word := "test"
		newDefinition := "new definition"
		err := dictionary.Update(word, newDefinition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)

	})

	t.Run("update a new word", func(t *testing.T) {
		newWord := "newWord"
		newDefinition := "new definition"
		err := dictionary.Update(newWord, newDefinition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, newWord, newDefinition)

	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word", err)
	}
	if definition != got {
		t.Errorf("got %q , want %q", got, definition)
	}
}
