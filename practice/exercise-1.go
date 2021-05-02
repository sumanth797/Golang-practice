package main

import "fmt"

const (
	A = iota
	B = iota
	C
	D = iota + 5
	E
)
const (
	A1 = iota
	B1 = iota
	C1
)

func main() {
	x, y, z := 42, "JamesBond", true
	fmt.Println(x, y, z)
	const (
		a int    = 2
		b string = "how"
	)
	fmt.Println(A, B, C, D, E)
	fmt.Println(A1, B1, C1)

}
