package repo

import (
	"database/sql"
	"fmt"

	"github.com/dropdevrahul/exchange-rates-go/internal/models"
)

type RateRepo interface {
	Create(d *models.DBAdapter, r *models.Rate) error
	List(d *models.DBAdapter) ([]models.Rate, error)
}

type Rates struct {
	TableName string
}

func (r *Rates) Create(d *models.DBAdapter, u *models.Rate) error {
	q := fmt.Sprintf(
		"Insert into %s (date, USD, EUR, GBP) VALUES ($1, $2, $3, $4)",
		r.TableName)
	_, err := d.DB.Exec(q, u.Date, u.USD, u.EUR, u.GBP)

	return handleError(err)
}

func (r *Rates) List(d *models.DBAdapter) ([]models.Rate, error) {
	res := []models.Rate{}
	q := fmt.Sprintf(
		"Select *  from %s order by date desc limit 10;",
		r.TableName)
	err := d.DB.Select(&res, q)

	return res, err
}

func handleError(err error) error {
	if err == nil {
		return nil
	}

	if err == sql.ErrNoRows {
		return models.ErrNotFound
	}

	return err
}
