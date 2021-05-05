package main

import "fmt"

func main() {
	// ans := variadic(1, 2, 3, 4, 5, 6)
	// fmt.Println(ans)
	defer fmt.Println("with defer")
	slice := []int{1, 3, 5, 7, 9}
	ans, a := variadic("hello", slice...)
	fmt.Println(ans, a)
	fmt.Println("without defer")

}

func variadic(str string, x ...int) (int, string) {
	fmt.Println(str)
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum, str

}
