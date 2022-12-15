package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := mux.Vars(r)["name"]
		switch name {
		case "bitcoin":
			response(w, "BTC")
		case "canadiandollar":
			response(w, "CAD")
		case "dogecoin":
			response(w, "DOGE")
		case "dollar":
			response(w, "USD")
		case "euro":
			response(w, "EUR")
		case "pound":
			response(w, "GBP")
		default:
			response(w, name)
		}
	})

	log.Print("The is Server Running on localhost port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
