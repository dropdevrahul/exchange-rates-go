package models

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

type DBAdapter struct {
	DB *sqlx.DB
}

var ErrNotFound = errors.New("error not found")
