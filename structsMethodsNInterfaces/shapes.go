package perimeter

import "math"

// import "fmt"

// Rectangle define the collection of fields used to store rectangle data
type Rectangle struct {
	Width float64
	Height float64
}
// Circle define the collection of fields used to store circle data
type Circle struct {
	Radius float64
}
// Triangle define the collection of fields used to store circle data
type Triangle struct {
	Base float64
	Height float64
}

// Shape defines methods to use on shapes structs
type Shape interface {
	Area() float64
}

// Perimeter recieve width and height and returns the perimeter of a rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Width)
}

// Area returns the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area returns the area of a circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Area returns the area of a triangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height)/2
}

