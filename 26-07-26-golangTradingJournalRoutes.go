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

func getTrader (w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Fprintln(w, "Only GET method is allowed here")
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		fmt.Fprintln(w, "Please provide a trader id, e.g. ?id=1")
		return
	}

	fmt.Fprintln(w, "You successfully requested trader id")
}

func createTrade (w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintln(w, "Only POST method is allowed here")
		return
	}

	var t Trade
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		fmt.Fprintln(w, "Error reading data", err)
		return
	}

	fmt.Fprintln(w, "Trade Entered", t.Asset, t.Direction, t.TraderID)
}

func getTrade(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Fprintln(w, "Only POST method is allowed here")
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		fmt.Fprintln(w, "Please provide a trade id, e.g. ?id=1")
		return
	}

	fmt.Fprintln(w, "You successfully requested trade id:", id)
}

func createTraderProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintln(w, "Only POST method is allowed here")
		return
	}

	var t TraderProfile
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		fmt.Fprintln(w, "Error reading data", err)
		return
	}

	fmt.Fprintln(w, "You successfully created a Trader Profile", t.PNL, t.WinRate, t.TraderID)
}

func main () {
	http.HandleFunc("/traders", createTrader)
	http.HandleFunc("/traders/get", getTrader)
	http.HandleFunc("/trades", createTrade)
	http.HandleFunc("/trades/get", getTrade)
	http.HandleFunc("/traderprofile", createTraderProfile)
	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}