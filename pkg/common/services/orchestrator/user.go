package orchestrator

import (
	"Cafeteria/pkg/common/models"
	"context"
)

type IUserOrchestrator interface {
	GetByName(ctx context.Context, user, password string) (*models.Users, error)
}
type UserOrchestrator struct {
	Engine *Engine
}

func NewUserOrchestrator(engine *Engine) IUserOrchestrator {
	return &UserOrchestrator{engine}
}

func (or *UserOrchestrator) GetByName(ctx context.Context, user, password string) (*models.Users, error) {
	dataBoard, err := or.Engine.DataStore.UserRepository.GetByName(ctx, user, password)
	if err != nil {
		return nil, err
	}
	return dataBoard, nil
}
