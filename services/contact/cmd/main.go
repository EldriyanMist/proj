package main

import (
	"log"
	"Assignment/pkg/store/postgres"
)

func main() {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "postgres"
	dbname := "Microservices"

	db, err := postgres.ConnectToDB(host, port, user, password, dbname)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}
