package main

import (
	"encoding/json"
	"fmt"
	"github.com/gosuri/uilive"
	"gopkg.in/alecthomas/kingpin.v2"
	"io/ioutil"
	"net/http"
	"time"
)

const ver = "0.1.0"

var (
	app     = kingpin.New("Keeper", "Crypto currency status tool.")
	version = kingpin.Flag("version", "Show current version of Keeper.").Short('v').Bool()

	currency  = kingpin.Flag("currency", "Change displayed currency").Short('c').Required().String()
	shortHelp = kingpin.CommandLine.HelpFlag.Short('h')
)

type Coin []struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
	Rank         string `json:"rank"`
	PriceUSD     string `json:"price_usd"`
	Price        string `json:"price_btc"`
	VolUSD24h    string `json:"24h_volume_usd"`
	MarketCapUSD string `json:"market_cap_usd"`
	Supply       string `json:"available_supply"`
	Total        string `json:"total_supply"`
	Max          string `json:"max_supply"`
	Change1h     string `json:"percent_change_1h"`
	Change24h    string `json:"percent_change_24h"`
	Change7d     string `json:"percent_change_7d"`
	LastUpdated  string `json:"last_updated"`
}

func getTop10Coin() string {
	var url string
	var all bool
	if *currency == "top" {
		url = fmt.Sprintf("https://api.coinmarketcap.com/v1/ticker/?limit=50")
		all = true
	} else {
		url = fmt.Sprintf("https://api.coinmarketcap.com/v1/ticker/%s", *currency)
		all = false
	}
	var coin Coin
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &coin)
	if err != nil {
		panic(err)
	}
	var header string
	if all {
		for x := 0; x < len(coin); x++ {
			header += fmt.Sprintf("Name: %v [%v]\nPrice (USD): $%s\nChange (24h): %v%%\n---\n",
				coin[x].Name,
				coin[x].Symbol,
				coin[x].PriceUSD,
				coin[x].Change24h)
		}
	} else {
		header += fmt.Sprintf("Name: %v [%v]\nPrice (USD): $%s\nChange (1h): %v%%\nChange (24h): %v%%\n",
			coin[0].Name,
			coin[0].Symbol,
			coin[0].PriceUSD,
			coin[0].Change1h,
			coin[0].Change24h)
	}
	return header
}

func checkFlags() {
	writer := uilive.New()
	writer.Start()
	kingpin.Parse()
	switch {
	case *version:
		fmt.Println(ver)
	case *currency != "":
		for {
			fmt.Fprintf(writer, "%v", getTop10Coin())
			time.Sleep(1 * time.Minute)
		}
	}
}

func main() {
	checkFlags()
}
