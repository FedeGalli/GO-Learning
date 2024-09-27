package arrayandslices

import "testing"

func TestAdd(t *testing.T) {
	got := Add([]int{1, 2, 3, 4})
	want := 10

	if got != want {
		t.Errorf("Want %v, got %v", want, got)
	}
}

func TestAddAll(t *testing.T) {
	got := AddAll([]int{1, 2, 3, 4}, []int{4})
	want := []int{10, 4}

	for i, num := range want {
		if num != got[i] {
			t.Errorf("Sum Want %v, got %v", want, got)
		}
	}
}
