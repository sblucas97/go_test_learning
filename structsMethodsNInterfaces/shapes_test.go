package perimeter

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
func TestArea(t *testing.T) {

	// anonymous struct definition
	areaTest := []struct {
		name 	string
		shape 	Shape
		hasArea 	float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 10.0, Height: 10.0}, hasArea: 100.0},
		{name: "Circle", shape: Circle{Radius: 10.0}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTest {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
	
			if got != tt.hasArea {
				t.Errorf("%#v got %f want %f", tt.shape, got, tt.hasArea)
			}
		})
	}

}

