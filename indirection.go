package main

import "fmt"

type Vertex7 struct {
	X, Y float64
}

func (v *Vertex7) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex7, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex7{3, 4}
	v.Scale(2) //go解释为(&v).Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex7{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
