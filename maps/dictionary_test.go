package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	word := "test"
	definition := "defition test"
	dictionary := Dictionary{ word: definition}

	t.Run("update definition", func(t *testing.T) {
		newDefinition := "new definition test"
		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("update definition of a unknow word", func(t *testing.T) {
		unknowWord := "unknownWord"
		newDefinition := "new definition test"
		err := dictionary.Update(unknowWord, newDefinition)

		assertError(t, err, ErrWordDoesNotExists)
		assertDefinition(t, dictionary, word, newDefinition)
	})
}

func TestDelete(t * testing.T) {
	word := "test"
	definition := "defition test"
	dictionary := Dictionary{ word: definition}

	t.Run("delete existing word", func(t *testing.T) {
		err := dictionary.Delete(word)

		assertError(t, err, nil)
	})

	t.Run("delete no existing word", func(t *testing.T) {
		noExistingWord := "car"

		err := dictionary.Delete(noExistingWord)

		assertError(t, err, ErrWordDoesNotExists)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search("test")

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got erro %q want %q", got, want)
	}
}
