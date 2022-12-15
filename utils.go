package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/perylemketwer/currency-api/models"
)

const Url = "https://economia.awesomeapi.com.br/json/last/"

func getCoinPrice(url string, coin string) string {
	resp, err := http.Get(url + coin)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var price models.CoinToReal
	json.Unmarshal([]byte(body), &price)
	switch coin {
	case "BTC":
		return price.BitcoinToReal.Bid
	case "CAD":
		return price.CanadianDolarToReal.Bid
	case "DOGE":
		return price.DogecoinToReal.Bid
	case "EUR":
		return price.EuroToReal.Bid
	case "GBP":
		return price.PoundToReal.Bid
	case "USD":
		return price.DolarToReal.Bid
	default:
		return "Coin not found!"
	}
}

func response(w http.ResponseWriter, coin string) {
	res := fmt.Sprint(getCoinPrice(Url, coin))
	if res == "Coin not found!" {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["price"] = res
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonResp)
}
