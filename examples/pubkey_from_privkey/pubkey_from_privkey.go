package main

import (
	"log"

	"github.com/funmi4194/go-bitcoin"
)

func main() {

	// create a private first to generte the address
	privateKey, err := bitcoin.CreatePrivateKeyString()
	if err != nil {
		log.Fatalf("error occurred: %s", err.Error())
	}

	// get the pubkey from privateKey string generated above.
	pubKey, err := bitcoin.PubKeyFromPrivateKeyString(privateKey)
	if err != nil {
		log.Fatalf("error occurred: %s", err.Error())
	}

	// Success!
	log.Printf("created pubkey: %s from private key: %s", pubKey, privateKey)
}
