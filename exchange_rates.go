package exchange

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// Currencies
const (
	USD string = "USD"
	EUR string = "EUR"
	GBP string = "GBP"
)

type Rate struct {
	USD float64
	EUR float64
	GBP float64
}

// last 10 days will be donated by string 'yyyy-dd-mm' as key and value as result
type ExchangeRateResult map[string]Rate

type ExchangeRateClient interface {
	// get exchange rates for the given date in yyyy-mm-dd format
	// base currency in relative of which exchange rates qre required
	Get(date string, base string) (ExchangeRateResult, error)
}

type ExchangeRatesAPIResponse struct {
	Success bool               `json:"success"`
	Date    string             `json:"date"`
	Base    string             `json:"base"`
	Rates   map[string]float64 `json:"rates"`
}

type ExchangeRatesAPI struct {
	ApiKey string
	Url    string
	client http.Client
}

func NewClient(client http.Client, apiKey, url string) *ExchangeRatesAPI {
	return &ExchangeRatesAPI{
		client: http.Client{},
		ApiKey: apiKey,
		Url:    url,
	}
}

// get exchange rates for the given date in yyyy-mm-dd format
func (e *ExchangeRatesAPI) Get(date string, base string) (res ExchangeRatesAPIResponse, err error) {
	url := fmt.Sprintf("%s/%s?access_key=%s&symbols=%s", e.Url, date, e.ApiKey,
		strings.Join([]string{USD, EUR, GBP}, ","))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return res, fmt.Errorf("Unable to create Request: %w ", err)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := e.client.Do(req)

	if err != nil {
		return res, fmt.Errorf("Unable to send request: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, fmt.Errorf("Unable to read response: %w", err)
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		return res, fmt.Errorf("Unable to parse response: %w", err)
	}

	return res, nil
}

func (e *ExchangeRatesAPI) GetLastTenDays(base string) (ExchangeRateResult, error) {
	ch := make(chan string, 10)
	results := make(chan ExchangeRatesAPIResponse, 10)
	response := ExchangeRateResult{}

	today := time.Now()

	go func(b string, c chan string, d chan ExchangeRatesAPIResponse) {
		for date := range c {
			res, err := e.Get(date, b)
			if err != nil {
				log.Println(err)
			}

			d <- res
		}
		close(d)
	}(base, ch, results)

	for i := 0; i < 10; i++ {
		date := today.AddDate(0, 0, -i).Format("2006-01-02")
		ch <- date
	}

	close(ch)
	for r := range results {
		response[r.Date] = Rate{
			USD: r.Rates[USD],
			EUR: r.Rates[EUR],
			GBP: r.Rates[GBP],
		}
	}
	return response, nil
}
