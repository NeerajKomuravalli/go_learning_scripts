package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

type Cryptoresponse []struct {
	Name              string `json:"name"`
	Price             string `json:"price"`
	Rank              string `json:"rank"`
	High              string `json:"high"`
	CirculatingSupply string `json:"circulating_supply"`
}

func (crResp Cryptoresponse) printMe() {
	fmt.Printf("Name : %s\nPrice : %s\nRank : %s\nHigh : %s\nCirculatingSupply : %s\n", crResp[0].Name, crResp[0].Price, crResp[0].Rank, crResp[0].High, crResp[0].CirculatingSupply)
}

func main() {
	// Can call it like : go run main.go -fiat=EUR -crypto=ETH
	fiatCurrency := flag.String(
		"fiat", "USD", "The name of the fiat currency you would like to know the price of your crypto in",
	)
	nameOfCrypto := flag.String(
		"crypto", "BTC", "Input the name of the CryptoCurrency you would like to know the price of",
	)
	flag.Parse()
	url := fmt.Sprintf("https://api.nomics.com/v1/currencies/ticker?key=3990ec554a414b59dd85d29b2286dd85&interval=1d&ids=%s&convert=%s", *nameOfCrypto, *fiatCurrency)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer resp.Body.Close()
	var cryptoresponse Cryptoresponse
	err = json.NewDecoder(resp.Body).Decode(&cryptoresponse)
	if err != nil {
		log.Fatalln(err)
	}
	cryptoresponse.printMe()
}
