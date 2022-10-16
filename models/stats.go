package models

import (
	"fmt"
)

type Stats struct {
	Rot    int `json:"rot"`
	Usages int `json:"usages"`
}

func GetStats(app *App) []Stats {
	schema := "SELECT * FROM report WHERE datetime > NOW() - INTERVAL '1 day';"
	rows, err := app.DB.Query(schema)
	if err != nil {
		fmt.Println(err)
	}

	tmp := make(map[int]int)
	for rows.Next() {
		// note that city can be NULL, so we use the NullString type
		var id int
		var rot int
		var datetime string
		rows.Scan(&id, &rot, &datetime)
		tmp[rot] += 1
	}
	// check the error from rows
	var res []Stats
	for rot, usages := range tmp {
		stats := Stats{
			Rot:    rot,
			Usages: usages,
		}
		res = append(res, stats)
	}

	return res
}
