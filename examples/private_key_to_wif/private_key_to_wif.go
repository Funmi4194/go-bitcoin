package main

import (
	"log"

	"github.com/funmi4194/go-bitcoin"
)

func main() {

	// Start with a private key (we will make one for this example)
	privateKey, err := bitcoin.CreatePrivateKeyString()
	if err != nil {
		log.Fatalf("error occurred: %s", err.Error())
	}

	// Create a wif
	var privateWif string
	if privateWif, err = bitcoin.PrivateKeyToWifString(privateKey, bitcoin.Mainnet); err != nil {
		log.Fatalf("error occurred: %s", err.Error())
	}

	// Success!
	log.Printf("private key: %s converted to wif: %s", privateKey, privateWif)
}
