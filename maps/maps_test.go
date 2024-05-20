package maps

import "testing"

func TestSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		dict := Dictionary{"test": "this is just a test"}

		got, err := dict.Search("test")
		want := "this is just a test"
		assertNoError(t, err)
		assertStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		dict := Dictionary{"test": "this is just a test"}

		_, err := dict.Search("unknown")
		want := ErrNotFound

		if err == nil {
			t.Fatal("expected to get an err")
		}
		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add("test", "this is just a test")
		assertNoError(t, err)
		want := "this is just a test"
		got, err := dictionary.Search("test")
		if err != nil {
			t.Fatal("should find added word:", err)
		}
		assertStrings(t, got, want)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		def := "this is a test"
		dict := Dictionary{word: def}
		err := dict.Add(word, "new def")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dict, word, def)
	})

}

func TestUpdate(t *testing.T) {
	t.Run("update an existing word", func(t *testing.T) {
		word := "test"
		def := "this is a test"

		d := Dictionary{word: def}
		newDef := "new def"
		err := d.Update(word, newDef)
		assertDefinition(t, d, word, newDef)
		assertNoError(t, err)
	})

	t.Run("update a new word", func(t *testing.T) {
		word := "test"
		def := "this is a test"

		d := Dictionary{}

		err := d.Update(word, def)
		assertError(t, err, ErrWordDoesNotExist)

	})
}

func TestDelete(t *testing.T) {
	t.Run("delete an existing word", func(t *testing.T) {
		word := "test"
		d := Dictionary{word: "this is a test"}

		d.Delete(word)
		_, err := d.Search(word)
		if err != ErrNotFound {
			t.Errorf("Expected %q to be deleted", word)
		}
	})

}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition)
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
