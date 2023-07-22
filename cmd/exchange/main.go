package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dropdevrahul/exchange-rates-go"
)

func main() {
	key := os.Getenv("API_KEY")
	client := exchange.NewClient(http.Client{}, key, "http://api.exchangeratesapi.io")

	res, err := client.GetLastTenDays(exchange.USD)
	if err != nil {
		log.Print(err)
	}

	for date, r := range res {
		fmt.Println(date)
		fmt.Printf("USD: %f\n", r.USD)
		fmt.Printf("EUR: %f\n", r.EUR)
		fmt.Printf("GBP: %f\n", r.GBP)
		fmt.Println("")
	}
}
