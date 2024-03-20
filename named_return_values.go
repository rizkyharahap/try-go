package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Rizki"
	middleName = "Mahfuddin"
	lastName = "Harahap"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}
