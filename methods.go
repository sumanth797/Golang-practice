package main

import "fmt"

type direction struct {
	speed    int
	duration string
}
type method interface {
	check()
}
type dd struct {
	speed    int
	duration string
}

//Receiver for struct
func (d direction) check() {
	fmt.Println("speed is ", d.speed, " and duration is ", d.duration)
}
func (de dd) check() {
	fmt.Println(de.duration)
}

//Interface receiver
func bar(m method) {
	//Assertion
	switch m.(type) {
	case direction:
		fmt.Println("value is ", m.(direction).duration)
	case dd:
		fmt.Println("speed is", m.(dd).speed)
	}
	fmt.Println(m)
}
func main() {
	out := direction{speed: 12, duration: "here"}
	in := dd{speed: 15, duration: "ten"}
	out.check()
	//Anonymous function
	sl := func(x int) int {
		fmt.Println("x is ", x)
		return x * 2
	}
	defer fmt.Println(bars()(), "is the answer")
	fmt.Println(sl(42))
	//Polymorphism ability
	bar(out)
	bar(in)
}
func bars() func() int {
	return func() int {
		return 555
	}
}
