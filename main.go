package main

import "fmt"

type Shape interface {
	Draw()
}

type Circle struct {
	radius float64
}

func (c *Circle) Draw() {
	fmt.Println("Menggambar lingkaran dengan radius", c.radius)
}

type Square struct {
	length float64
}

func (s *Square) Draw() {
	fmt.Println("Menggambar persegi dengan panjang", s.length)
}

func ShapeFactory(shapeType string) Shape {
	switch shapeType {
	case "circle":
		return &Circle{radius: 5.0}
	case "square":
		return &Square{length: 10.0}
	default:
		return nil
	}
}

func main() {
	circle := ShapeFactory("circle")
	circle.Draw()

	square := ShapeFactory("square")
	square.Draw()
}