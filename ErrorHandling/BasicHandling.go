package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}
	return a / b, nil

}

func idiomatic() {
	result, err := divide(20, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Success:", result)
	}
}

func errorf(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Invalid input:denominator %.2f cannot be zero", b)
	}
	return a / b, nil
}

func readConfig() error {
	return fmt.Errorf("Config error:%w", errors.New("File not found"))
}

func main() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result)

	idiomatic()

	result1, err1 := errorf(10, 2)
	if err1 != nil {
		fmt.Println("Error1:", err1)
		return
	}
	fmt.Println("Result2:", result1)

	readConfig()

}
