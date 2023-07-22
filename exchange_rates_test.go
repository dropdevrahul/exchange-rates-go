package exchange_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dropdevrahul/exchange-rates-go"
	"github.com/stretchr/testify/assert"
)

func TestGetSuccess(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"success":true,"historical":true,"date":"2013-01-24","timestamp":1387929599,"base":"USD","rates":{"USD":1.000000,"EUR":1.196476,"GBP":1.739516}}`))
	}))

	e := exchange.NewClient(http.Client{}, "key", s.URL)
	res, err := e.Get("2023-01-24", "USD")

	exp := exchange.ExchangeRatesAPIResponse{
		Success: true,
		Date:    "2013-01-24",
		Rates: map[string]float64{
			"USD": 1.000000,
			"EUR": 1.196476,
			"GBP": 1.739516,
		},
		Base: "USD",
	}
	assert.Equal(t, nil, err)
	assert.Equal(t, exp, res)
}
