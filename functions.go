package main

import "fmt"

func main() {
	defer fmt.Println("with defer")
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ans, a := variadic("hello", slice...)
	fmt.Println(ans, a)
	fmt.Println("without defer")
	addition := even(add, slice...)
	fmt.Println("Addition of even slice is ", addition)
	multiplication := even(multiply, slice...)
	fmt.Println("Multiplication of even slice is ", multiplication)

}

func variadic(str string, x ...int) (int, string) {
	fmt.Println(str)
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum, str

}

// Callback Function
func even(f func(va ...int) int, slice ...int) int {
	var s1 []int
	for _, v := range slice {
		if v%2 == 0 {
			s1 = append(s1, v)
		}
	}
	return f(s1...)
}
func add(v ...int) int {
	sum := 0
	for _, v := range v {
		sum += v
	}
	return sum
}
func multiply(v ...int) int {
	result := 1
	for _, v := range v {
		result *= v
	}
	return result
}
