package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := r.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			msg := fmt.Sprintf("no such item: %v\n", item)
			http.Error(w, msg, http.StatusNotFound)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		msg := fmt.Sprintf("no such page: %s\n", r.URL)
		http.Error(w, msg, http.StatusNotFound)
	}
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}
