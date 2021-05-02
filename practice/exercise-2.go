package main

import (
	"fmt"
)


var x int = 2
var y string = "Bond"
var z bool = true

type hotdog int

var h hotdog

func main() {
	s := fmt.Sprintf("%v,%v,%v", x, y, z)
	fmt.Printf("%v,%v", s, h)
	fmt.Println()
	fmt.Println(runtime.)
}
