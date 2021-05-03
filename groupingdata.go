package main

import "fmt"

func main() {
	var a [5]int
	a[0] = 14
	a[4] = 24
	fmt.Println(a)
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b[4])
}
