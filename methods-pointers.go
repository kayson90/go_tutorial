package main

import (
	"fmt"
	"math"
)

type Vertex5 struct {
	X, Y float64
}

func (v Vertex5) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//指针接收者可以改变值，值接收者传递的是副本
func (v *Vertex5) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex5{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
