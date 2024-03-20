package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	rizki := Man{"Rizki"}
	rizki.Married()

	fmt.Println(rizki.Name)
}
