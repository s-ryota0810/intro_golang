package main

import "fmt"

func main() {

	// for i := 0; i < 10; i++ {
	// 	if i == 3 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// i := 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	//無限ループ
	i := 0
	for {
		fmt.Println(i)
		i++
		if i == 40 {
			break
		}
	}
}
