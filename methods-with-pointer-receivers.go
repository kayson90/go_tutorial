package main

import (
	"fmt"
	"math"
)

//使用指针接收器的原因，1.需要修改 2.性能更高
type Vertex9 struct {
	X, Y float64
}

func (v *Vertex9) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//虽然没有改变值，也使用了指针接收器
func (v *Vertex9) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex9{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
