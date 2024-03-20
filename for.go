package main

import "fmt"

func regularForLoop() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	fmt.Println("Selesai")
}

func rangeForLoop() {
	names := []string{"Rizki", "Mahfuddin", "Harahap"}
	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	}
}

func breakForLoop() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println("Perulangan ke", i)
	}
}

func continueForLoop() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("Perulangan ke", i)
	}
}

func main() {
	regularForLoop()
	rangeForLoop()
	breakForLoop()
	continueForLoop()
}
