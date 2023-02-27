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

	Client := snipcart.NewClient(*snipcartApiKey)

	updateOrder := snipcart.SnipcartOrderUpdate{
		Status: snipcart.Delivered,
	}

	response, err := Client.UpdateOrder("b35990df-c0ca-4014-94de-1caa7bd7bb51", &updateOrder)
	if err != nil {
		log.Fatal(err)
	}

	for k, v := range response.Items {
		log.Printf("%v: %v\n", k, v)
	}
}
