package main

import (
	"log"

	"github.com/funmi4194/go-bitcoin"
)

func main() {

	// Convert the wif into a private key
	privateKey, err := bitcoin.WifToPrivateKey("5K4psRpsyqZmioyQ3wwxm17N7e1HbDLx2j2nn3NcmwfH166hgQj")
	if err != nil {
		log.Fatalf("error occurred: %s", err.Error())
	}

	// Get an address
	var address string
	if address, err = bitcoin.GetAddressFromPrivateKey(privateKey, bitcoin.NativeSegwit, bitcoin.Mainnet); err != nil {
		log.Fatalf("error occurred: %s", err.Error())
	}

	// Success!
	log.Printf("found address: %s from private key: %s", address, privateKey)
}
