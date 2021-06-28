package main

/*
func hi(name string) string { //戻り値の型を指定する
	// fmt.Println("hi!" + name)
	msg := "hi!" + name
	return msg
}

func hi(name string) (msg string) {
	// fmt.Println("hi!" + name)
	msg = "hi!" + name
	return
}
func main() {
	//hi("taguchi")
	fmt.Println(hi("inoe"))
}
*/

import "fmt"

func main() {
	//fmt.Println(swap(5, 2))
	f := func(a, b int) (int, int) {
		return b, a
	}
	fmt.Println(f(2, 3))
	
	//即時関数
	func(msg string) {
		fmt.Println(msg)
	}("taguchi")
}
