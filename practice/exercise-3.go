package main

import "fmt"

const (
	c int = 4
	b     = 5
)
const (
	y1 = 2021 - iota
	y2
	y3
	y4
)

func main() {
	// fmt.Printf("%v", c+b)
	a := 16
	fmt.Printf("%v <-> %b <-> %#x", a, a, a)
	b := a << 1
	fmt.Printf("\n%v <-> %b <-> %#x\n", b, b, b)
	////Raw String literal
	s := `Here is the 
	example of a raw string literal
	and it is "just fine"
	and
	end`
	fmt.Println(s)
	fmt.Println(y1, y2, y3, y4)
}
