package main

import "fmt"

func main() {
	rectangle1 := newRectangle(10, 10)

	fmt.Println(rectangle1.getName())

	data := &SurfaceCalculator{
		surface: 0,
	}

	rectangle1.accept(data)

	fmt.Println(data.surface)

	circle1 := newCircle(3)

	fmt.Println(circle1.getName())

	square1 := newSquare(31)

	fmt.Println(square1.getName())

}
