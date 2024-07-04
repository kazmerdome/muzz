package explore_test

import (
	"testing"

	"github.com/kazmerdome/muzz/internal/module/explore"
	explore_grpc "github.com/kazmerdome/muzz/internal/module/explore/explore-grpc"
	"github.com/stretchr/testify/assert"
)

func TestModule(t *testing.T) {
	exploreModule := explore.NewExploreModule(nil)
	decisionController := exploreModule.GetController()
	assert.Implements(t, (*explore_grpc.ExploreServiceServer)(nil), decisionController)
}
