package integers

import "testing"

func TestAdder(t *testing.T) {
	got := Add(1, 2)
	want := 3

	if got != want {
		t.Errorf("Wanted %v, Got %v", want, got)
	}
}
