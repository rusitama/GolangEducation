package main

import (
	"fmt"
	"math"
)

// Интерфейс Shape
type Shape interface {
	Area() float64
}

// Структура Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Метод Area для Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Структура Circle
type Circle struct {
	Radius float64
}

// Метод Area для Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func main() {
	// Пример использования
	rect := Rectangle{Width: 5, Height: 3}
	circ := Circle{Radius: 2}

	// Встраиваем Shape в Rectangle и Circle
	var rectShape Shape = rect
	var circShape Shape = circ

	// Вывод площадей
	fmt.Println("Площадь прямоугольника:", rectShape.Area())
	fmt.Println("Площадь круга:", circShape.Area())
}
