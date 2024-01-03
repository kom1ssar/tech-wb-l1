package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

// расчитываем дистанцию двух точек
func (p *Point) Distance(p2 *Point) float64 {

	dy := p2.y - p.y
	dx := p2.x - p.x

	return math.Sqrt(
		math.Pow(float64(dx), 2) +
			math.Pow(float64(dy), 2))
}

func (p *Point) GetX() float64 {
	return p.x
}

func (p *Point) GetY() float64 {
	return p.y
}

func (p *Point) SetX(x float64) {
	p.x = x
}

func (p *Point) SetY(y float64) {
	p.y = y
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func main() {

	firstPoint := NewPoint(5, 10)
	secondPoint := NewPoint(10, 3)
	fmt.Println(firstPoint.Distance(secondPoint))
}
