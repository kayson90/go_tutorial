package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	//结构体文法
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	q  = &Vertex{1, 2}
)

func main() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	//结构体指针
	p := &v
	p.X = 1e9
	fmt.Println(v)

	fmt.Println(v1, q, v2, v3)

}
