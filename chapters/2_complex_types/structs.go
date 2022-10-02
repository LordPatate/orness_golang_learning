package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	p := &v
	p.X = 1e9 // no need for (*p).X
	fmt.Println(v)

	origin := &Vertex{}     // X = 0 and Y = 0
	x_base := &Vertex{X: 1} // Y = 0
	y_base := &Vertex{Y: 1} // X = 0

	fmt.Println(origin, x_base, y_base)
}
