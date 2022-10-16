package models

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type EnvData struct {
	IP         string
	Port       string
	PGPassword string
	DBName     string
	User       string
	URL        string
}

func LoadEnvData() *EnvData {
	fmt.Println("Reading env file")
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	fmt.Println("Env file read successfuly")

	var data EnvData
	data.IP = os.Getenv("IP")
	data.Port = os.Getenv("PORT")
	data.DBName = os.Getenv("DBNAME")
	data.PGPassword = os.Getenv("PGQLPASSWORD")
	data.User = os.Getenv("DBUSER")

	data.URL = fmt.Sprintf("postgres://%v:%v@%v:%v/%v", data.User, data.PGPassword, data.IP, data.Port, data.DBName) // postgres://testuser:testpassword@localhost/testmooc
	fmt.Println(data.URL)

	return &data
}
