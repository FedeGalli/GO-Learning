package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	got, ok := dictionary.Search("test")
	want := "this is just a test"
	if got != want || ok != nil {
		t.Errorf("Got %q, want %q", got, want)
	}
}
