package main

import (
	"log"

	"github.com/funmi4194/go-bitcoin"
)

func main() {
	// Start with a private key (we will make one for this example)
	privateKey, err := bitcoin.CreatePrivateKey()
	if err != nil {
		log.Fatalf("error occurred: %s", err.Error())
	}

	// Get an address
	var address string
	if address, err = bitcoin.GetAddressFromPrivateKey(privateKey, bitcoin.Taproot, bitcoin.Mainnet); err != nil {
		log.Fatalf("error occurred: %s", err.Error())
	}

	// Get the script
	var script string
	if script, err = bitcoin.GetScriptFromAddress(address, bitcoin.Mainnet); err != nil {
		log.Fatalf("error occurred: %s", err.Error())
	}

	// Success!
	log.Printf("generated script: %s from address: %s", script, address)
}
