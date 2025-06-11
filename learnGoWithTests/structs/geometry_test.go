package structs

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	expected := 40.0

	if got != expected {
		t.Errorf("got %.2f expected %.2f", got, expected)
	}
}

func TestAreaForStructs(t *testing.T) {
	t.Run("Test rectangle area", func(t *testing.T) {
		r := Rectangle{5.0, 5.0}
		expected := 25.0

		AreaCheck(t, r, expected)
	})
	t.Run("Test circle area", func(t *testing.T) {
		c := Circle{5.0}
		expected := 78.5

		AreaCheck(t, c, expected)
	})
}

func AreaCheck(t testing.TB, shape Shape, expected float64) {
	t.Helper()
	got := shape.Area()
	const epsilon = 0.0001
	if math.Abs(got-expected) > epsilon {
		t.Errorf("%#v got %.2f expected %.2f", shape, got, expected)
	}
}

// Table driven tests
func TestAreas(t *testing.T) {
	areasToTest := []struct {
		shape    Shape
		expected float64
	}{
		{shape: Rectangle{12, 6}, expected: 72.0},
		{shape: Circle{10}, expected: 314.0},
		{shape: Triangle{12, 6}, expected: 36.0},
	}

	for _, tt := range areasToTest {
		AreaCheck(t, tt.shape, tt.expected)
	}
}
