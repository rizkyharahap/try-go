package main

import "fmt"

func random() any {
	return "OK"
}

func main() {
	// var result any = random()
	// var resultString string = result.(string)

	// fmt.Println(resultString)

	// var resultInt int = result.(int) // panic
	// fmt.Println(resultInt)

	result := random()

	switch result.(type) {
	case string:
		fmt.Println("String", result)
	case int:
		fmt.Println("Int", result)
	default:
		fmt.Println("Unknown", result)
	}
}
