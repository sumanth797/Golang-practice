package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("False")

	case !true:
		fmt.Println("True")

	case (2 == 2):
		fmt.Println("true")
		fallthrough
	case (3 != 3):
		fmt.Println("prints")
	}
}
