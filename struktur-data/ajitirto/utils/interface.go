package utils

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func Interface() {
	// case 1
	rectangle := Rectangle{Width: 10, Height: 20}
	fmt.Println("Nilai area 20 x 10 = ", rectangle.Area())

	// case 2
	var shape Shape = Rectangle{Width: 10, Height: 20}
	fmt.Println("Nilai area 20 x 10 = ", shape.Area())

}