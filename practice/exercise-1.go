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
const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
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
	fmt.Printf("%v--%b", kb, kb)
	fmt.Printf("\n%v--%b", mb, mb)
	fmt.Printf("\n%v--%b", gb, gb)

}
