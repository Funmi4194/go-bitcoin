package main

import (
	"log"

	"github.com/funmi4194/go-bitcoin"
)

func main() {

	// Create a wif
	wifString, err := bitcoin.CreateWifString(bitcoin.Mainnet)
	if err != nil {
		log.Fatalf("error occurred: %s", err.Error())
	}

	// Success!
	log.Printf("wif key: %s", wifString)
}
