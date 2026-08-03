package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Trader struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

type Trade struct {
	ID int `json:"id"`
	Asset string `json:"asset"`
	EntryPrice int `json:"entry_price"`
	ExitPrice int `json:"exit_price"`
	Direction string `json:"direction"`
	TraderID int `json:"trader_id"`
}

type TraderProfile struct {
	ID int `json:"id"`
	TraderID int `json:"trader_id"`
	PNL int `json:"pnl"`
	WinRate int `json:"win_rate"`
}

var traders []Trader
var trades []Trade
var traderProfiles []TraderProfile

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

	traders = append(traders, t)
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
	
	idNum, err := strconv.Atoi(id)
	if err != nil {
		fmt.Fprintln(w, "Invalid id format")
		return
	}

	for _, trader := range traders {
		if trader.ID == idNum {
			fmt.Fprintln(w, "Found trader:", trader.Name, trader.Email)
			return
		}
	}

	fmt.Fprintln(w, "Trader not found")
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

	trades = append(trades, t)
	fmt.Fprintln(w, "Trade Entered", t.Asset, t.Direction, t.TraderID)
}

func getTrade(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Fprintln(w, "Only GET method is allowed here")
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		fmt.Fprintln(w, "Please provide a trade id, e.g. ?id=1")
		return
	}

	idNum, err := strconv.Atoi(id)
	if err != nil {
		fmt.Fprintln(w, "Invalid id format")
		return
	}

	for _, trade := range trades {
		if trade.ID == idNum {
			fmt.Fprintln(w, "Found trade:", trade.Asset, trade.Direction, trade.TraderID)
			return
		}
	}

	fmt.Fprintln(w, "Trade not found")
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

	traderProfiles = append(traderProfiles, t)
	fmt.Fprintln(w, "You successfully created a Trader Profile", t.PNL, t.WinRate, t.TraderID)
}

func getTraderProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Fprintln(w, "Only GET method is allowed here")
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		fmt.Fprintln(w, "Please provide a trader profile id, e.g. ?id=1")
		return
	}

	idNum, err := strconv.Atoi(id)
	if err != nil {
		fmt.Fprintln(w, "Invalid id format")
		return
	}

	for _, traderProfile := range traderProfiles {
		if traderProfile.ID == idNum {
			fmt.Fprintln(w, "Found Profile:", traderProfile.PNL, traderProfile.WinRate, traderProfile.TraderID)
			return
		}
	}

	fmt.Fprintln(w, "Profile not found")
}

func main () {
	http.HandleFunc("/traders", createTrader)
	http.HandleFunc("/traders/get", getTrader)
	http.HandleFunc("/trades", createTrade)
	http.HandleFunc("/trades/get", getTrade)
	http.HandleFunc("/traderprofile", createTraderProfile)
	http.HandleFunc("/traderprofile/get", getTraderProfile)
	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}