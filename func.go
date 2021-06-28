package main

import "fmt"

/*
func hi(name string) string { //戻り値の型を指定する
	// fmt.Println("hi!" + name)
	msg := "hi!" + name
	return msg
}
*/

func hi(name string) (msg string) {
	// fmt.Println("hi!" + name)
	msg = "hi!" + name
	return
}
func main() {
	//hi("taguchi")
	fmt.Println(hi("inoe"))
}
