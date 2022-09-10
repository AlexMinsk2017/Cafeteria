package datastore

import (
	"github.com/elastic/go-elasticsearch/v8"
)

type DataElastic struct {
	elHandler *elasticsearch.Client
}

func NewElasticDataStore(elHandler *elasticsearch.Client) *DataElastic {
	elasticDataStore := DataElastic{
		elHandler: elHandler,
	}
	return &elasticDataStore
}
