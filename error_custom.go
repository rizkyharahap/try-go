package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"Vallidation Error"}
	}

	if id != "123" {
		return &notFoundError{"Data Not Found"}
	}

	return nil
}

func main() {
	err := SaveData("123", nil)

	if err != nil {
		// if validationErr, ok := err.(*validationError); ok {
		// 	fmt.Println("validation error:", validationErr.Error())
		// } else if notFoundError, ok := err.(*notFoundError); ok {
		// 	fmt.Println("not found error:", notFoundError.Error())
		// } else {
		// 	fmt.Println("unknown error:", err.Error())
		// }

		switch err.(type) {
		case *validationError:
			fmt.Println("validation error:", err.Error())
		case *notFoundError:
			fmt.Println("not found error:", err.Error())
		default:
			fmt.Println("unknown error:", err.Error())
		}
	} else {
		fmt.Println("Success")
	}

}
