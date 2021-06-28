/*
string "hello"
int 		53
float64	10.2
bool		true
nil

デフォルト
var s string // ""
var a int			// 0
var f bool		// false

*/

package main

import "fmt"

func main() {
	fmt.Println() //改行つき
	a := 10
	b := 12.3
	c := "hoge"
	var d bool
	fmt.Printf("a: %d, b:%f, c:%s, d:%t\n", a, b, c, d)
	fmt.Printf("a: %v, b:%v, c:%v, d:%v\n", a, b, c, d)
}
