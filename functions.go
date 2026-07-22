package main

import (
	"fmt"
	"strconv"
)

func identification(name string) string {
	return "What is your name? " + name + ":)"
}

func favorite(food string) string {
	return "What is your favorite food? " + food + "!"
}

func big(a int, b int) string {
	if a > b {
		return "Think of two numbers in your head right now and between those numbers, tell me which is the biggest! " + strconv.Itoa(a)
	} else {
		return "Think of two numbers in your head right now and between those numbers, tell me which is the biggest! " + strconv.Itoa(b)
	}
}

func full(name string, food string, big int) string {
	return "So your name is " + name + ", your favorite food is " + food + ", and your biggest number is " + strconv.Itoa(big)
}

func main() {
	message := identification("Ghazi")
    fmt.Println(message)

	favs := favorite("Pizza")
	fmt.Println(favs)

    result := big(6, 7)
    fmt.Println(result)

	sentence := full("Ghazi" , "Pizza", 7)
	fmt.Println(sentence)
}