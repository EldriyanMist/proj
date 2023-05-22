package main

import (
	"log"

	"github.com/your/package/postgres"
)

func main() {
	host := "localhost"
	port := "5432"
	user := ""
	password := ""
	dbname := ""

	db, err := postgres.ConnectToDB(host, port, user, password, dbname)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}
