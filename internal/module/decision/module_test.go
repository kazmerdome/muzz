package decision_test

import (
	"testing"

	"github.com/kazmerdome/muzz/internal/module/decision"
	"github.com/kazmerdome/muzz/mocks"
	"github.com/stretchr/testify/assert"
)

type moduleFixture struct {
	mocks struct {
		db    *mocks.DB
		sqlDb *mocks.SqlDB
	}
}

func newModuleFixture(t *testing.T) *moduleFixture {
	f := &moduleFixture{}
	f.mocks.db = mocks.NewDB(t)
	f.mocks.sqlDb = mocks.NewSqlDB(t)
	return f
}

func TestModule(t *testing.T) {
	f := newModuleFixture(t)
	f.mocks.db.EXPECT().GetDB().Return(f.mocks.sqlDb)
	decisionModule := decision.NewDecisionModule(f.mocks.db)
	decisionRepo := decisionModule.GetRepository()
	assert.Implements(t, (*decision.DecisionRepository)(nil), decisionRepo)
}
