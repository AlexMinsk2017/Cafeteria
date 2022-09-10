package orchestrator

import "Cafeteria/pkg/common/services/datastore/infra/datastore"

type Engine struct {
	DataStore    *datastore.DataStore
	Orchestrator *Orchestrator
	ElasticData  *datastore.DataElastic
}
type Orchestrator struct {
	Engine           *Engine
	UserOrchestrator IUserOrchestrator
}

func NewOrchestrator(engine *Engine) *Orchestrator {
	newOrchestrator := Orchestrator{
		Engine:           engine,
		UserOrchestrator: NewUserOrchestrator(engine),
	}
	return &newOrchestrator
}
