# README

Keeper is a Cryptocurrency status tool. It uses the [CoinMarketCap](https://coinmarketcap.com) API and updates every minute.

## Requirements:

* Go 1.9.2 or higher

## Installation:

`go get -u github.com/kyoto-shift/keeper`

## Use:

`-h, --help` Basically displays this.  
`-v, --version` Prints the current version of Keeper.  
`-c, --currency` Changes the tracked Cryptocurrency. If set to `top`, it will track the top 50 Cryptocurrencies (according to CoinMarketCap).  

## Examples:

```
[user@computer]$ keeper -c bitcoin
Name: Bitcoin [BTC]
Price (USD): $15335.6
Change (1h): 0.49%
Change (24h): 1.44%
Change (7d): 5.76%
```

```
[user@computer]$ keeper -c top
Name: Bitcoin [BTC]
Price (USD): $15277.5
Change (24h): 1.04%
---
Name: Ripple [XRP]
Price (USD): $3.48639
Change (24h): 10.2%
---
Name: Ethereum [ETH]
Price (USD): $995.801
Change (24h): 5.21%
---
Name: Bitcoin Cash [BCH]
Price (USD): $2415.63
Change (24h): -7.33%
---
Name: Cardano [ADA]
Price (USD): $1.19986
Change (24h): 13.4%
---
...
```
