package main

import "fmt"

func main() {
	a := 5
	var pa *int
	pa = &a // &a = aのアドレスをさす
	// paの領域にあるデータの値 = *pa
	fmt.Println(*pa)
	fmt.Println(pa)
	fmt.Println(a)
	fmt.Println(&a)
	
}
