package orchestrator

type IUserOrchestrator interface {
}
type UserOrchestrator struct {
	Engine *Engine
}

func NewUserOrchestrator(engine *Engine) IUserOrchestrator {
	return &UserOrchestrator{engine}
}
