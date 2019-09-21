package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p *Point) Length() float64 {
	return math.Hypot((*p).X, (*p).Y)
}

func (p *Point) MoveXTo(newX float64) {
	(*p).X = newX
}

// by value
func (p Point) MoveXToByValue(newX float64) {
	p.X = newX
}

func (p *Point) MoveYTo(newY float64) {
	(*p).Y = newY
}

func main() {
	p := Point{3, 4}
	fmt.Println(p)
	fmt.Println(p.Length()) // 5
	p.MoveXTo(6)
	p.MoveXToByValue(56)
	fmt.Println("-----------------")
	fmt.Println(p.Length()) // will not change
}
