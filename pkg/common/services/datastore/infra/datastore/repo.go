package datastore

import (
	"Cafeteria/pkg/common/services/datastore/repo"
	"gorm.io/gorm"
)

type DataStore struct {
	dbHandler      *gorm.DB
	UserRepository repo.IUserRepository
}

func NewDataStore(dbHandler *gorm.DB) *DataStore {
	dataStore := DataStore{
		dbHandler:      dbHandler,
		UserRepository: repo.NewUser(dbHandler),
	}
	return &dataStore
}

func (ds *DataStore) WithTransaction(fn func(tx *gorm.DB) error) error {
	return ds.dbHandler.Transaction(fn)
}
