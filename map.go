package main

import "fmt"

func main() {
	// create variable with type
	// var person map[string]string = map[string]string{}
	// person["name"] = "Rizki"
	// person["address"] = "Kendal"

	person := map[string]string{
		"name":    "Rizki",
		"address": "Kendal",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	fmt.Println(person["salah"])

	book := make(map[string]string)

	book["title"] = "Buku Golang"
	book["author"] = "Rizki"
	book["ups"] = "Salah"

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)
}
