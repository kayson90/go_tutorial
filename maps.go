package main

import "fmt"

type Vertex1 struct {
	Lat, Long float64
}

var m map[string]Vertex1

func main() {
	m = make(map[string]Vertex1)
	m["Bell Labs"] = Vertex1{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
