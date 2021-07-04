package main

import (
	"fmt"
	"math"
)

type Vertex4 struct {
	X, Y float64
}

//方法和函数并没有本质区别，只不过方法有接收器
func Abs(v Vertex4) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex4{3, 4}
	fmt.Println(Abs(v))
}
