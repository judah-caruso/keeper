# README

Keeper is a Cryptocurrency status tool using the [CoinMarketCap](https://coinmarketcap.com) API.

## Requirements:

* Go 1.9.2 or higher

## Installation:

`go get -u github.com/kyoto-shift/keeper`

## Use:

`-h, --help` Basically displays this.  
`-v, --version` Prints the current version of Keeper.  
`-c, --currency` Changes the tracked Cryptocurrency. If set to `all`, Keeper will track the top 10 (unless changed by `-a`) Cryptocurrencies (according to CoinMarketCap).  
`-a --amount` Sets the number of currencies `-c all` will track. If `-c all` isn't used, this will have no effect.

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
[user@computer]$ keeper -c all
Name: Bitcoin [BTC] (1)
Price (USD): $15277.5
Change (24h): 1.04%
---
Name: Ripple [XRP] (2)
Price (USD): $3.48639
Change (24h): 10.2%
---
Name: Ethereum [ETH] (3)
Price (USD): $995.801
Change (24h): 5.21%
---
Name: Bitcoin Cash [BCH] (4)
Price (USD): $2415.63
Change (24h): -7.33%
---
Name: Cardano [ADA] (5)
Price (USD): $1.19986
Change (24h): 13.4%
---
...
```

# CHANGELOG

**0.2.0**  
* Added `-i, --interval` flag which changes the tracker's update time.
* Added `-a, --amount` flag which changes the number of top currencies tracked.
* Added rank of currencies if multiple are being tracked.
* Changed `top` to `all` if `-c` is used.
* Changed default value of `all` from 50 to 10. It can now also be changed with `-a [number]`

**0.1.2**  
* Initial release.
