package main

import (
	"fmt"
	"math"
)

type Vertex6 struct {
	X, Y float64
}

func Abs1(v Vertex6) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//作为参数和作为接收者，指针都是一样的道理
func Scale(v *Vertex6, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex6{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs1(v))
}
