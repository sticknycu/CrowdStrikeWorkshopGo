package main

import (
	"math"
)

type Shape interface {
	getName() string

	accept(visitor ShapeVisitor)
}

type ShapeVisitor interface {
	visitForSquare(square *Square)

	visitForRectangle(rectangle *Rectangle)

	visitForCircle(circle *Circle)
}

type SurfaceCalculator struct {
	surface float64
}

type CenterCoordinates struct {
	x float64
	y float64
}

func (rectangle *Rectangle) accept(visitor ShapeVisitor) {
	visitor.visitForRectangle(rectangle)
}

func (rectangle *Rectangle) getName() string {
	return "Rectangle"
}

func (square *Square) accept(visitor ShapeVisitor) {
	visitor.visitForSquare(square)
}

func (square *Square) getName() string {
	return "Square"
}

func (circle *Circle) accept(visitor ShapeVisitor) {
	visitor.visitForCircle(circle)
}

func (circle *Circle) getName() string {
	return "Circle"
}

func (calculator *SurfaceCalculator) visitForSquare(square *Square) {
	calculator.surface = square.side * square.side
}

func (calculator *SurfaceCalculator) visitForRectangle(rectangle *Rectangle) {
	calculator.surface = rectangle.width * rectangle.length
}

func (calculator *SurfaceCalculator) visitForCircle(circle *Circle) {
	calculator.surface = math.Pow(circle.radius, 2) * math.Pi
}

type Square struct {
	side float64
}

func newSquare(newSide float64) *Square {
	return &Square{side: newSide}
}

type Rectangle struct {
	width  float64
	length float64
}

func newRectangle(newWidth float64, newLength float64) *Rectangle {
	return &Rectangle{width: newWidth, length: newLength}
}

type Circle struct {
	radius float64
}

func newCircle(newRadius float64) *Circle {
	return &Circle{radius: newRadius}
}
