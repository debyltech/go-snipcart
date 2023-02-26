package main

import (
	"flag"
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

	response, err := snipcartProvider.GetOrdersByStatus(snipcart.Processed)
	if err != nil {
		log.Fatal(err)
	}

	for k, v := range response.Items {
		log.Printf("%v: %v\n", k, v)
	}
}
