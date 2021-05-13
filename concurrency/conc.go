package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("CPU's are ", runtime.NumCPU())
	wg.Add(1)
	go foo()
	go bar()
	fmt.Println("Go routines are ", runtime.NumGoroutine())
	wg.Wait()
	//time.Sleep(time.Millisecond)
}
func bar() {
	for i := 0; i < 5; i++ {
		fmt.Println("Bar - ", i)
		wg.Done()
	}
}
func foo() {
	for i := 0; i < 5; i++ {
		fmt.Println("Foo - ", i)
		wg.Done()
	}
}

//wg.Done()
