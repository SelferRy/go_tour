package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vertex) Square() (float64, float64) {
	return math.Pow(v.X, 2), math.Pow(v.Y, 2)
}

func (v Vertex) Perimeter() float64 {
	return v.X + v.Y
}

func (v Vertex) Some(s string) string {
	return s
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(v.Square())
	fmt.Println(v.Perimeter())
	fmt.Println(v.Some("the string"))
}
