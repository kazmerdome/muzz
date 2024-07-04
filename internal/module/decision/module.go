package decision

import (
	"github.com/kazmerdome/muzz/internal/actor/db"
	decisionQuerier "github.com/kazmerdome/muzz/internal/module/decision/decision-querier"
)

//go:generate make name=DecisionRepository mock
//go:generate make name=Querier structname=DecisionQuerier filename=DecisionQuerier.go srcpkg=github.com/kazmerdome/muzz/internal/module/decision/decision-querier mock

type decisionModule struct {
	repository DecisionRepository
}

func NewDecisionModule(db db.DB) *decisionModule {
	querier := decisionQuerier.New(db.GetDB())
	repository := NewDecisionRepository(querier)
	return &decisionModule{
		repository: repository,
	}
}

func (m *decisionModule) GetRepository() DecisionRepository {
	return m.repository
}
