package main

import (
	"fmt"
	"math"
)

//不一定要是结构体作为接收器，都可以作为接收器
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
