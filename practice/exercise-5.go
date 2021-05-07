package main

import "fmt"

func main() {
	fmt.Println("exercise")
	a := 5
	p := &a
	fmt.Println(*p)
	ptr(p)
	fmt.Println(a)
}
func ptr(v *int) {
	*v = 6
}
