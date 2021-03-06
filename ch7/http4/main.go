package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.list)
	http.HandleFunc("/update", db.update)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type database map[string]dollars

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

func (db database) list(w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		http.Error(w, fmt.Sprintf("no such item: %q\n", item), http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) update(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price_s := r.URL.Query().Get("price")
	price, err := strconv.ParseFloat(price_s, 32)
	if err != nil {
		http.Error(w, fmt.Sprintf("Wrong price format: %q\n", price), http.StatusBadRequest)
		return
	}
	db[item] = dollars(price)
	fmt.Fprintf(w, "%s: $%.2f\n", item, price)
}
