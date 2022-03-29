package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	//declaring a slice of struct using anon. w/ 2 fields and fill it with cases
	areaTests := []struct {
		shape Shape
		hasArea float64
	}{
		//easier to understand case fields  
		{shape: Rectangle{12, 6}, hasArea: 72.0},
		{shape: Circle{10}, hasArea: 314.1592653589793},
		{shape: Triangle{12, 6}, hasArea: 36.0},

		//this is still acceptable
		// {Rectangle{12, 6}, 72.0},
		// {Circle{10}, 314.1592653589793},
		// {Triangle{12, 6}, 36.0},
	}
	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
		}
	}
}