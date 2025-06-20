package day_05

import (
	"fmt"
	"math"
)

// interface : collection of signature of methods

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Length  float64
	Breadth float64
}

// to implement shape interface, rectangle struct must implement all the methods in the interface

func (r Rectangle) Area() float64 {
	return r.Length * r.Breadth
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Breadth)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius

}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// polymorphism concept

func DisplayShapeInfo(shape Shape) {
	fmt.Printf("Area: %.2f\n", shape.Area())
	fmt.Printf("Perimeter: %.2f\n", shape.Perimeter())
}

func CheckInterface() {

	cir := Circle{Radius: 5}
	rect := Rectangle{Length: 5, Breadth: 10}
	fmt.Println("Circle = ", cir)
	DisplayShapeInfo(cir)
	fmt.Println()
	fmt.Println("Rectangle = ", rect)
	DisplayShapeInfo(rect)
}
