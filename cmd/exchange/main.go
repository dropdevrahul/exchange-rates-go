package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dropdevrahul/exchange-rates-go"
	app "github.com/dropdevrahul/exchange-rates-go/internal"
	"github.com/dropdevrahul/exchange-rates-go/internal/models"
)

func main() {
	key := os.Getenv("API_KEY")

	app := app.NewApp()

	client := exchange.NewClient(http.Client{}, key, "http://api.exchangeratesapi.io")
	res, err := client.GetLastTenDays(exchange.USD)
	if err != nil {
		log.Print(err)
	}

	for date, r := range res {
		t, err := time.Parse("2006-01-02", date)
		if err != nil {
			log.Println(err)
		}

		dbRate := models.Rate{
			EUR:  r.EUR,
			USD:  r.USD,
			GBP:  r.GBP,
			Date: t,
		}

		// we can do this concurrently but that requires wait groups
		err = app.Repos.Create(app.DB, &dbRate)
		if err != nil {
			log.Println(err)
		}
	}

	dbList, err := app.Repos.List(app.DB)
	if err != nil {
		log.Println(err)
	}
	for _, res := range dbList {
		fmt.Println(res.Date)
		fmt.Printf("USD: %f\n", res.USD)
		fmt.Printf("EUR: %f\n", res.EUR)
		fmt.Printf("GBP: %f\n", res.GBP)
		fmt.Println("")
	}
}
