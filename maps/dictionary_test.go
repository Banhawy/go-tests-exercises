package dictionary

import "testing"

func TestSearch(t *testing.T) {
	// dictionary := map[string]string{"test": "this is just a test"}
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {

		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(got, want, t)
	})

	t.Run("unkown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertError(err, ErrWordNotFound, t)

	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		assertError(err, nil, t)
		assertDefinition(dictionary, t, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, definition)

		assertError(err, ErrWordExists, t)
		assertDefinition(dictionary, t, definition)

	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		dictionary.Update(word, newDefinition)

		assertDefinition(dictionary, t, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(err, ErrWordDoesNotExist, t)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	assertError(err, ErrWordNotFound, t)
}

func assertDefinition(dictionary Dictionary, t *testing.T, definition string) {
	t.Helper()

	got, err := dictionary.Search("test")
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(got, definition, t)
}

func assertStrings(got string, want string, t *testing.T) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q, given %q", got, want, "test")
	}
}

func assertError(got, want error, t *testing.T) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q, given %q", got, want, "test")
	}
}
