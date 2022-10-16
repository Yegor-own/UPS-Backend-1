package main

import (
	"fmt"
	"net/http"
	"ups/models"
)

var app = models.NewApp()

func init() {
	result := app.DB.MustExec("CREATE TABLE report (Id SERIAL NOT NULL PRIMARY KEY, Rot INTEGER, Datetime TIMESTAMPTZ NOT NULL DEFAULT NOW());")
	fmt.Println("Table created", result)
}

func main() {
	app.Router.HandleFunc("/encode", EncodeController).Methods("POST")
	app.Router.HandleFunc("/decode", DecodeController).Methods("GET")
	app.Router.HandleFunc("/stats", StatsController).Methods("GET")
	http.ListenAndServe(":3000", app.Router)
}
