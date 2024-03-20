package main

import (
	"fmt"
	"try-go/database"
	_ "try-go/internal"
)

func main() {
	fmt.Println(database.GetDatabase())
}
