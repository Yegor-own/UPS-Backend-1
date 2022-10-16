package models

import (
	"fmt"
)

type RotReport struct {
	Id       int    `db:"Id"`
	Rot      int    `db:"Rot"`
	Datetime string `db:"Datetime"`
}

func RotReportWrite(rot int, app *App) {
	res, err := app.DB.Exec(fmt.Sprintf("INSERT INTO report (rot) VALUES (%v);", rot))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
