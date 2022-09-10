package repo

import "gorm.io/gorm"

type IUserRepository interface {
}
type UserRepository struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}
