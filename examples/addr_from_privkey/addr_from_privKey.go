package main

import (
	"log"

	"github.com/funmi4194/go-bitcoin"
)

func main() {

	// create a private first to generte the address
	privateKey, err := bitcoin.CreatePrivateKey()
	if err != nil {
		log.Fatalf("error occurred: %s", err.Error())
	}

	// get the address from privateKey generated above.
	// make sure to specify the network type and address type you want
	// for mainnet network type pass bitcoin.Mainnet
	// for NativeSegwit type pass bitcoin.NativeSegwit
	address, err := bitcoin.GetAddressFromPrivateKey(privateKey, bitcoin.NativeSegwit, bitcoin.Mainnet)
	if err != nil {
		log.Fatalf("error occurred: %s", err.Error())
	}

	// Success!
	log.Printf("found address: %s from private key: %s", address, privateKey)
}
