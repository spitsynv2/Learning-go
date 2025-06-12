package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "testing"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		expected := "testing"

		assertStrings(t, got, expected)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	t.Run("new word", func(t *testing.T) {
		word := "test"
		expected := "testing"

		err := dictionary.Add("test", "testing")

		assertError(t, err, nil)
		assertMapContains(t, dictionary, word, expected)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		expected := "testing"

		err := dictionary.Add("test", "testing")

		assertError(t, err, ErrContainsSameElement)
		assertMapContains(t, dictionary, word, expected)
	})
}

func TestUpdate(t *testing.T) {
	dictionary := Dictionary{"test": "testing"}
	t.Run("existing word", func(t *testing.T) {
		word, expected := "test", "testing2"
		err := dictionary.Update(word, "testing2")

		assertError(t, err, nil)
		assertMapContains(t, dictionary, word, expected)
	})
	t.Run("not existing word", func(t *testing.T) {
		err := dictionary.Update("test2", "testing2")

		assertError(t, err, ErrNotFound)
	})
}

func TestDelete(t *testing.T) {
	dictionary := Dictionary{"test": "testing"}
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		err := dictionary.Delete(word)

		assertError(t, err, nil)

		_, err = dictionary.Search(word)

		assertError(t, err, ErrNotFound)
	})
	t.Run("not existing word", func(t *testing.T) {
		err := dictionary.Delete("test2")

		assertError(t, err, ErrNotFound)
	})
}

func assertStrings(t testing.TB, got, expected string) {
	t.Helper()

	if got != expected {
		t.Errorf("got %q want %q", got, expected)
	}
}

func assertError(t testing.TB, got, expected error) {
	t.Helper()

	if got != expected {
		t.Errorf("got error %q want %q", got, expected)
	}
}

func assertMapContains(t testing.TB, d Dictionary, word, expected string) {
	t.Helper()

	got, err := d.Search(word)

	if err != nil {
		t.Fatal("Should find added word:", err)
	}
	assertStrings(t, got, expected)
}
