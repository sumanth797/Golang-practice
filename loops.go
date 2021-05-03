package main

import "fmt"

func main() {
	fmt.Println("For Loop")
	x := 1
	for {

		if x > 5 {
			break
		}
		fmt.Println(x)
		x++
	}
	fmt.Println("Done")

	for a := 35; a < 50; a++ {
		fmt.Printf("%v----%c", a, a)
		fmt.Println()
	}

	if x := 23; x == 23 {
		fmt.Println(x, x == 23)
	}
	fmt.Println("DONE")
	fmt.Println("END")
}
