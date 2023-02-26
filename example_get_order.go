package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/debyltech/go-snipcart/snipcart"
)

func main() {
	snipcartApiKey := flag.String("key", "", "Snipcart API Key")
	flag.Parse()

	if *snipcartApiKey == "" {
		log.Fatal("missing -key flag")
	}

	snipcartProvider := snipcart.NewSnipcartProvider(*snipcartApiKey)

	response, err := snipcartProvider.GetOrder("b35990df-c0ca-4014-94de-1caa7bd7bb51")
	if err != nil {
		log.Fatal(err)
	}

	byteResponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(byteResponse))
}
