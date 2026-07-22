package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	Name string
	Age int
}

func (p Person) Greet() string {
	return "Hi, I'm " + p.Name
}

func (p Person) Numero() string {
	return "And i am " + strconv.Itoa(p.Age)
}

func main() {
	p := Person{Name: "Ghazi", Age: 20}
	fmt.Println(p.Name, p.Age)
	fmt.Println(p.Greet())
	fmt.Println(p.Numero())
}