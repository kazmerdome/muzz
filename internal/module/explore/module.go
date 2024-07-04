package explore

import (
	"github.com/kazmerdome/muzz/internal/module/decision"
)

//go:generate make name=ExploreService mock

type exploreModule struct {
	controller *exploreController
}

func NewExploreModule(decisionRepository decision.DecisionRepository) *exploreModule {
	service := NewExploreService(decisionRepository)
	controller := NewExploreController(service)

	return &exploreModule{
		controller: controller,
	}
}

func (m *exploreModule) GetController() *exploreController {
	return m.controller
}
