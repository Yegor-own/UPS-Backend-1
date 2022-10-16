package models

import (
	"fmt"
	"log"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sqlx.DB
}

func NewApp() *App {
	envData := LoadEnvData()

	fmt.Println("Connecting to DB")
	db, err := sqlx.Connect("postgres",
		"host="+envData.IP+
			" user="+envData.User+
			" password="+envData.PGPassword+
			" dbname="+envData.DBName+
			" sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("DB connected successful")

	if err = db.Ping(); err != nil {
		log.Fatalln(err)
	}

	return &App{
		Router: mux.NewRouter(),
		DB:     db,
	}
}
