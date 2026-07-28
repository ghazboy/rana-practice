package main

import (
	"fmt"
)

type Trader struct {
	ID int
	Name string
	Email string
}

type Trade struct {
	ID int
	Asset string
	EntryPrice int
	ExitPrice int
	Direction string
	TraderID int
}

type TraderProfile struct {
	ID int
	TraderID int
	PNL int
	WinRate int
}

func main() {
	t := Trader{ID: 1, Name: "Ghazi", Email: "ghazi.khairy@gmail.com"}
	fmt.Println(t)
}