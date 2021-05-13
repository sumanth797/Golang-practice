package main

import "fmt"

type person struct {
	firstname string
	lastname  string
}

//Embedded Struct
type degree struct {
	person    //....Unqualified type name acts as field name
	firstname int
	course    string
}
type foo int

var f foo

const fool int = 42

func main() {
	f = 42 + foo(fool)
	fmt.Printf("%T,%v\n", f, f)
	p1 := person{firstname: "sai", lastname: "venkat"}
	p2 := person{firstname: "Sri", lastname: "sri"}
	c1 := degree{person{"ant", "man"}, 297, "mtech"}
	fmt.Println(p1.firstname, p2.lastname)
	fmt.Println(c1.person.firstname)
	fmt.Println(c1.firstname)
	//...Anonymous structs....
	a := struct {
		number int
		name   string
	}{
		12, "sum",
	}
	fmt.Println(a)

}
