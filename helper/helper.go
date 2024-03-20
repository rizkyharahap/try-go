package helper

import "fmt"

var version = "1.0.0"
var Application = "golang" // dapat diakses diluar package

func SayHello(name string) string {
	return "Hello " + name
}

func sayGoogBye(name string) string {
	return "Good Bye " + name
}

func Contoh() {
	fmt.Println(sayGoogBye("Rizki"))
	fmt.Println(version)
}
