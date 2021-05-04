package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	flavours  []string
	num       string
}

func main() {
	p1 := person{"sum", "div", []string{"choco", "vanila", "berry"}, "school"}
	p2 := person{"ang", "el", []string{"idk", "cone", "bc"}, "col"}
	p3 := person{"eng", "al", []string{"kidk", "cone"}, "lcol"}
	// p := []person{p1, p2}
	m := map[string]person{
		p1.firstname: p1,
		p2.firstname: p2,
		p3.firstname: p3,
	}
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v.firstname)
		for k1, v1 := range v.flavours {
			fmt.Println(k1, v1)
		}
		fmt.Println("-----------------")

	}

	// fmt.Println("fname", p1.firstname)
	// for k, v1 := range p1.flavours {
	// 	fmt.Println(k, v1)
	// }
	// fmt.Println("fname", p2.firstname)
	// for k, v1 := range p2.flavours {
	// 	fmt.Println(k, v1)
	// }
}
