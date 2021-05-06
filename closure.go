package main

import "fmt"

func main() {
	fmt.Println("Closure")
	a := Closure()
	b := Closure()
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(a())
	fmt.Println(b())

}
func Closure() func() int {
	var x int
	return func() int {
		x++
		return x

	}
}
