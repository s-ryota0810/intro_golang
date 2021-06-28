package main

import "fmt"

type greeter interface {
	greet()
}

type japanese struct {
}

type american struct {
}

func (j japanese) greet() {
	fmt.Println("こんにちは")
}

func (a american) greet() {
	fmt.Println("Hello")
}

// interface{}
func show(t interface{}) {
	// 型アサーション
	// _, ok := t.(japanese)
	// if ok {
	// 	fmt.Println("私は日本人です")
	// } else {
	// 	fmt.Println("I am not japanese")
	// }
	// 型Switch
	switch t.(type) {
	case japanese:
		fmt.Println("私は日本人です")
	default:
		fmt.Println("I am not a Japanese")
	}

}

func main() {
	greeters := []greeter{japanese{}, american{}}
	for _, greeter := range greeters {
		greeter.greet()
		show(greeter)
	}
}
