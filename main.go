package main

import "fmt"

///go fmt ./...
//creating own type
type owntype int

var o owntype

func main() {
	// Examples()
	// sqrt()
	a := 2
	// fmt.Println(Second() * 2 * 4)
	o = 45
	fmt.Println(a, a+int(o)+1)
}
