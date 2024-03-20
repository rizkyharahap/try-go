package main

import (
	"fmt"
	"try-go/helper"
)

func main() {
	result := helper.SayHello("Rizki")
	fmt.Println(result)

	fmt.Println(helper.Application)
	// fmt.Println(helper.version)             // tidak bisa diakses
	// fmt.Println(helper.sayGoodBye("Rizki")) // tidak bisa diakses
}
