package main

import (
	"Cafeteria/pkg/common/config"
	"Cafeteria/pkg/common/db"
	"Cafeteria/pkg/common/services/datastore/infra/datastore"
	"log"
)

func main() {
	con, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	client := db.Init(&con)
	dataStore := datastore.NewDataStore(client)
}
