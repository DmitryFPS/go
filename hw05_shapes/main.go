package main

import (
	"fmt"
	"math"
)

type Shape interface {
	CalculateArea() float64
}

type Circle struct {
	Radius float64
}

func (circle *Circle) CalculateArea() float64 {
	return math.Pi * math.Pow(circle.Radius, 2)
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (rectangle *Rectangle) CalculateArea() float64 {
	return rectangle.Width * rectangle.Height
}

type Triangle struct {
	Base   float64
	Height float64
}

func (triangle *Triangle) CalculateArea() float64 {
	return 0.5 * triangle.Base * triangle.Height
}

func calculateArea(anyShape any) (float64, error) {
	shape, ok := anyShape.(Shape)
	if ok {
		return shape.CalculateArea(), nil
	}
	return 0, fmt.Errorf("переданный объект не является фигурой: %s", anyShape)
}

func main() {
	circle := &Circle{5}
	rectangle := &Rectangle{10, 5}
	triangle := &Triangle{8, 6}

	shapes := []any{circle, rectangle, triangle, "Что то совсем не понятное :)"}

	for i := 0; i < len(shapes); i++ {
		area, err := calculateArea(shapes[i])
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(area)
		}
	}
}
