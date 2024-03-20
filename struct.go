package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int8
}

/**
* Method sayHello
**/
func (customer Customer) sayHello() {
	fmt.Println("Hello, My name is", customer.Name)
}

func strucBasic() {
	var customer Customer
	customer.Name = "Rizki Mahfuddin Harahap"
	customer.Address = "Dimana yah"
	customer.Age = 100

	fmt.Println(customer)
}

func structLiterals() {
	customer := Customer{
		Name:    "Rizki Mahfuddin Harahap",
		Address: "Dimana yah",
		Age:     100,
	}

	fmt.Println(customer)
}

func main() {
	strucBasic()
	structLiterals()

	kampang := Customer{Name: "Kampang"}
	kampang.sayHello()
}
