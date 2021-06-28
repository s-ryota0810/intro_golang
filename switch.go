package main

import "fmt"

func main() {
	//signal := "blue"

	// switch signal {
	// case "red":
	// 	fmt.Println("stop")
	// case "Green", "blue":
	// 	fmt.Println("go")
	// case "yellow":
	// 	fmt.Println("caution")
	// default:
	// 	fmt.Println("wrong signal")
	// }
	score := 82
	switch {
	case score > 80:
		fmt.Println("Great!")
	case score > 60:
		fmt.Println("normal")
	default:
		fmt.Println("soso")
	}
}
