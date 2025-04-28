package main

import (
	"log"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/funmi4194/go-bitcoin"
)

func main() {
	// create a wif
	wifString, err := bitcoin.CreateWifString(bitcoin.Mainnet)
	if err != nil {
		log.Fatalf("error occurred: %s", err.Error())
	}

	// create a wif from a string
	var wifKey *btcutil.WIF
	wifKey, err = bitcoin.WifFromString(wifString)
	if err != nil {
		log.Fatalf("error occurred: %s", err.Error())
	}

	// success!
	log.Printf("wif key: %s is also: %s", wifString, wifKey.String())
}
