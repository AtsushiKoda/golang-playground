package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func structPointers() {
	v := Vertex{1, 2}
	p := &v
	p.X = 0x10 // (*p).xと書かなくてよい
	fmt.Println(v)
}
