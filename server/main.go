package main

import (
	"Cafeteria/pkg/common/config"
	"Cafeteria/pkg/common/db"
	"Cafeteria/pkg/common/services/datastore/infra/datastore"
	"Cafeteria/pkg/common/services/io/elastic"
	"Cafeteria/pkg/common/services/io/web"
	"Cafeteria/pkg/common/services/orchestrator"
	"log"
)

func main() {
	con, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	//postgres
	client := db.Init(&con)
	dataStore := datastore.NewDataStore(client)

	//elastic
	elasticClient, err := elastic.ClientElastic()
	if err != nil {
		log.Fatalln(err)
		return
	}
	elasticDataStore := datastore.NewElasticDataStore(elasticClient)

	//engine
	engine := orchestrator.Engine{
		DataStore:   dataStore,
		ElasticData: elasticDataStore,
	}
	engine.Orchestrator = orchestrator.NewOrchestrator(&engine)

	//webservice
	server := &web.WebServices{
		Orchestrator: engine.Orchestrator,
	}
	err = server.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
