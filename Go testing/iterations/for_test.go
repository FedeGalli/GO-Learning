package iterations

import (
	"ciccio"
	"testing"
)

func TestFor(t *testing.T) {
	got := Repeat("a", 1)
	want := "a"
	dictionary := ciccio.Dictionary{"test": "this is just a test"}

	dictionary.Add()

	if want != got {
		t.Errorf("Want %v, got %v", want, got)
	}
}

func BenchmarkFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 100000)
	}
}
