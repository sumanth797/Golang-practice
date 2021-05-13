package main

import "fmt"

func main() {
	// var a [5]int
	// a[0] = 14
	// a[1] = 24
	// a[2] = 34
	// a[3] = 44
	// a[4] = 54
	// for k, v := range a {
	// 	fmt.Printf("%v // %v // %T\n", k, v, v)
	// }
	b := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// for k, v := range b {
	// 	fmt.Println(k, v)
	// }
	c := append(b, 11, 12, 13)
	fmt.Println(c)
	x := []int{14, 15, 16}
	c = append(c, x...)
	fmt.Println(c)
	c = append(c[0:4], c[5:8]...)
	fmt.Println(c)
	s := make([]string, 5, 6)
	s = []string{"hello", "can", "u", "run", "it"}
	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}
	m1 := []string{"ab", "cd", "ef"}
	m2 := []string{"gh", "ij", "kl"}
	md := [][]string{m1, m2}
	for k, v := range md {
		fmt.Println(k, v)
		for k, j := range v {
			fmt.Println(k, j)
		}
	}
	m := map[string][]string{
		"sum": []string{"cricket", "football"},
		"cc":  []string{"play", "enjoy"},
	}
	m["gf"] = []string{"store", "phone"}
	fmt.Println(len(m))
	delete(m, "cc")
	fmt.Println(len(m))
	for k, v := range m {
		fmt.Println("for key", k)
		for ke, in := range v {
			fmt.Println(ke, in)
		}
	}
}
