package main

import "fmt"

func main() {
	// m := make(map[string]int) //[key]value
	// m["taguchi"] = 200
	// m["fkoji"] = 400
	// fmt.Println(m["taguchi"])
	m := map[string]int{"taguchi": 100, "fjiko": 200}
	fmt.Println(m)
	fmt.Println(len(m))
	delete(m, "taguchi")
	fmt.Println(m)
	v, Ok := m["fffjiko"]
	fmt.Println(v)
	fmt.Println(Ok)
}
