package main

import "fmt"

func greet(name string) string {
    return "Hello, " + name + "!"
}

func add(a int, b int) int {
    return a + b
}

func main() {
    message := greet("Ghazi")
    fmt.Println(message)

    result := add(5, 3)
    fmt.Println(result)
}