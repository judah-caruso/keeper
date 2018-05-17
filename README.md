# README

Keeper is a cryptocurrency status tool using the [CoinMarketCap](https://coinmarketcap.com) API.

## Requirements:

* Go 1.10.1 or higher

## Installation:

`go get -u github.com/kyoto-shift/keeper`

## Use:

`-h, --help` Displays this.  
`-v, --version` Prints the current version of Keeper.  
`-c, --currency` Changes the tracked cryptocurrency. If set to `all`, Keeper will track the top 10 (unless changed by `-n`) cryptocurrencies (according to CoinMarketCap).  
`-n --number` Sets the number of currencies `-c all` will track. If `-c all` isn't used, this will have no effect.

Set a default currency to track with:
```
[user@computer]$ export KEEPER_FAVORITE=Ethereum
```
To set a favorite permanently, add it to your bash_profile (or equivalent).

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

```
[user@computer]$ keeper -c all -n 3
Name: Bitcoin [BTC] (1)                                                   
Price (USD): $8068.52                                
Change (1h): -1.9%                                  
Change (24h): -3.09%                                     
---                                   
Name: Ethereum [ETH] (2)                                                    
Price (USD): $673.196                               
Change (1h): -2.15%                                  
Change (24h): -3.96%                                    
---                                   
Name: Ripple [XRP] (3)                                                    
Price (USD): $0.664309                                 
Change (1h): -1.53%                                 
Change (24h): -4.88%                                    
---
```

# CHANGELOG

**0.2.2**
* Keeper now looks for a favorite currency using the environment variable `KEEPER_FAVORITE`.
* Changed the flag `-a, --amount` to `-n, --number`.

**0.2.0**  
* Added `-i, --interval` flag which changes the tracker's update time.
* Added `-a, --amount` flag which changes the number of top currencies tracked.
* Added rank of currencies if multiple are being tracked.
* Changed `top` to `all` if `-c` is used.
* Changed default value of `all` from 50 to 10. It can now also be changed with `-a [number]`
* Fixed version flag being unusable unless `-c [crypto]` was used first.

**0.1.2**  
* Initial release.
