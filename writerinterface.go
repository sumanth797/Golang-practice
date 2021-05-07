package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	fmt.Println("hello")
	fmt.Fprintln(os.Stdout, "hello")
	io.WriteString(os.Stdout, "hello\n")
	//sort
	a := []int{2, 1, 4, 3, 6, 5, 8, 8, 7}
	fmt.Println(a)
	sort.Ints(a)
	fmt.Println(a)
}
