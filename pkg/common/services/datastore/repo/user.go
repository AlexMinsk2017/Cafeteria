package repo

import (
	"Cafeteria/pkg/common/models"
	"context"
	"gorm.io/gorm"
)

type IUserRepository interface {
	WithTx(tx *gorm.DB) IUserRepository
	GetByName(ctx context.Context, user, password string) (*models.Users, error)
}
type UserRepository struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (rep *UserRepository) WithTx(tx *gorm.DB) IUserRepository {
	return NewUser(tx)
}
func (rep *UserRepository) GetByName(ctx context.Context, user, password string) (*models.Users, error) {
	data := models.Users{}
	err := rep.db.WithContext(ctx).First(&data, "user = ?", user, "password = ?", password).Error
	return &data, err
}
