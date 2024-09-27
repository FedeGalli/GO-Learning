package interafces

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("Rectangle Area", func(t *testing.T) {
		got := Rectangle{10.0, 10.0}.Perimeter()
		want := 40.0

		if want != got {
			t.Errorf("Want %v, got %v", want, got)
		}
	})
	t.Run("Rectangle Area", func(t *testing.T) {
		got := Circle{10.0}.Perimeter()
		want := 62.83185307179586

		if want != got {
			t.Errorf("Want %v, got %v", want, got)
		}
	})
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("Want %g, got %g", want, got)
		}
	}
	t.Run("Rectangle Area", func(t *testing.T) {
		want := 100.0
		checkArea(t, Rectangle{10.0, 10.0}, want)
	})
	t.Run("Circle Area", func(t *testing.T) {
		want := 314.1592653589793
		checkArea(t, Circle{10.0}, want)
	})
}
