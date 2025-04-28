package main

import (
	"log"

	"github.com/funmi4194/go-bitcoin"
)

func main() {

	// Convert the wif into a private key
	privateKey, err := bitcoin.WifToPrivateKeyString("5K4psRpsyqZmioyQ3wwxm17N7e1HbDLx2j2nn3NcmwfH166hgQj")
	if err != nil {
		log.Fatalf("error occurred: %s", err.Error())
	}

	// Success!
	log.Printf("private key: %s", privateKey)
}
