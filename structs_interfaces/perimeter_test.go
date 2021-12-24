package structs_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10.0, 10.0}, 40.0},
		{Circle{5}, 31.41592653589793},
	}

	for _, tc := range perimeterTests {
		got, _ := tc.shape.Perimeter()
		if got != tc.want {
			// print out struct first when testing interface
			t.Errorf("%#v got %g want %g", tc.shape, got, tc.want)
		}
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10.0, 10.0}, 100.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tc := range areaTests {
		got, _ := tc.shape.Area()
		if got != tc.want {
			// print out struct first when testing interface
			t.Errorf("%#v got %g want %g", tc.shape, got, tc.want)
		}
	}
}
