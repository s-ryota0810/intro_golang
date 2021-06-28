package main

import "fmt"

func main() {
	b := []int{1, 2, 3, 23, 43}
	s := b[1:4] // [3]

	fmt.Println(s)
	fmt.Println(len(b), cap(b))
	fmt.Println(len(s), cap(s))
	fmt.Println(s)
	s[1] = 33
	fmt.Println(s)
	fmt.Println(len(s), cap(s))
}
