package main

import (
	"fmt"
	"errors"
)

func checkNumber(a int) (string, error) {
	if a < 0 {
		return "", errors.New("Please enter a positive number")
	}

	if a % 2 == 0 {
		return "Even", nil
	} else {
		return "Odd", nil
	}
}

func main() {
	result, err := checkNumber(7)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	result2, err2 := checkNumber(-67)
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Println("Result:", result2)
	}
}