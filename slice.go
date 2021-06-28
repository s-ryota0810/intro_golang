package main

import "fmt"

func main() {
	// s := make([]int, 3) // [0 0 0]
	s := []int{1, 3, 5}
	// append
	s = append(s, 5, 3, 565)
	//copy
	t := make([]int, len(s))
	n := copy(t, s)
	fmt.Println(s)
	fmt.Println(t)
	fmt.Println(n)
}
