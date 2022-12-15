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

func getCoinPrice(url string, coin string) (string, error) {
	resp, err := http.Get(url + coin)
	if err != nil {
		log.Fatal(err)
	}

	// Logging status code
	log.Printf("Status Code: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var price models.CoinToReal
	json.Unmarshal([]byte(body), &price)
	switch coin {
	case "BTC":
		return price.BitcoinToReal.Bid, nil
	case "CAD":
		return price.CanadianDolarToReal.Bid, nil
	case "DOGE":
		return price.DogecoinToReal.Bid, nil
	case "EUR":
		return price.EuroToReal.Bid, nil
	case "GBP":
		return price.PoundToReal.Bid, nil
	case "USD":
		return price.DolarToReal.Bid, nil
	default:
		return "Coin not found!", nil
	}
}

func response(w http.ResponseWriter, coin string) {
	res, err := getCoinPrice(Url, coin)
	if err != nil {
		fmt.Println(err)
	}

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
