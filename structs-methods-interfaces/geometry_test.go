package geometry

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{5.0, 10.0}
	got := Perimeter(rectangle)
	want := 30.0

	if got != want {
		t.Errorf("got %2f want %2f", got, want)
	}
}

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Area()
		want := 100.0

		assertCorrectCalculation(got, want, t)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		assertCorrectCalculation(got, want, t)
	})
}

func assertCorrectCalculation(got float64, want float64, t *testing.T) {
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}
