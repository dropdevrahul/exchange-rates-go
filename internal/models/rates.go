package models

import "time"

type Rate struct {
	ID   string    `db:"id"`
	USD  float64   `db:"usd"`
	EUR  float64   `db:"eur"`
	GBP  float64   `db:"gbp"`
	Date time.Time `db:"date"`
}
