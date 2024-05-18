package structs

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	tcs := []struct {
		name string
		r    Rectangle
		p    float64
	}{
		{"valid inputs", Rectangle{10.0, 10.0}, 40.0},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			act := tc.r.Perimeter()
			if act != tc.p {
				t.Errorf("got %.2f, expected %.2f, given %+v", act, tc.p, tc.r)
			}
		})
	}
}

func TestArea(t *testing.T) {
	tcs := []struct {
		name string
		s    Shape
		a    float64
	}{
		{"rectangle", &Rectangle{10.0, 10.0}, 100.0},
		{"circle", &Circle{10.0}, math.Pi * 100.0},
		{"triangle", &Triangle{12, 6}, 36.0},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			act := tc.s.Area()
			if act != tc.a {
				t.Errorf("got %g, expected %g, given %+v", act, tc.a, tc.s)
			}
		})
	}

}
