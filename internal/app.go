package app

import (
	"fmt"
	"log"
	"os"

	"github.com/dropdevrahul/exchange-rates-go/internal/models"
	"github.com/dropdevrahul/exchange-rates-go/internal/repo"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type App struct {
	DB       *models.DBAdapter
	Repos    repo.RateRepo
	settings Settings
}

func NewApp() *App {
	a := App{}
	a.LoadSettings()
	a.LoadDB()
	a.LoadRepos()

	return &a
}

type Settings struct {
	DB DBSettings `mapstructure:"db"`
}

type DBSettings struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

func (a *App) LoadSettings() {
	settings := Settings{}
	settings.DB.Host = os.Getenv("APP_DB_HOST")
	settings.DB.Port = os.Getenv("APP_DB_PORT")
	settings.DB.User = os.Getenv("APP_DB_USER")
	settings.DB.Password = os.Getenv("APP_DB_PASSWORD")
	settings.DB.Database = os.Getenv("APP_DB_DB")

	a.settings = settings
}

func (a *App) LoadRepos() {
	rateRepo := repo.Rates{
		TableName: "rates",
	}

	a.Repos = &rateRepo
}

func (a *App) LoadDB() error {
	log.Printf("connecting to db")

	db, err := sqlx.Connect("postgres",
		fmt.Sprintf("host=%s port =%s user=%s password=%s dbname=%s sslmode=disable",
			a.settings.DB.Host, a.settings.DB.Port, a.settings.DB.User,
			a.settings.DB.Password,
			a.settings.DB.Database,
		))

	if err != nil {
		log.Fatal("Failed connecting to database, ", err)
		return err
	}

	log.Printf("connected to db")
	a.DB = &models.DBAdapter{
		DB: db,
	}

	return nil
}
