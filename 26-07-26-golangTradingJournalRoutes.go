package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func createTrader (w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintln(w, "Only POST method is allowed here")
		return
	}

	var t Trader
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		fmt.Fprintln(w, "Error reading data", err)
		return
	}

	fmt.Fprintln(w, "Trader Created", t.Name, t.Email)
}

func main () {
	http.HandleFunc("/traders", createTrader)
	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}