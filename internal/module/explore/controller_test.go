package explore_test

import (
	context "context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/kazmerdome/muzz/internal/module/explore"
	explore_grpc "github.com/kazmerdome/muzz/internal/module/explore/explore-grpc"
	"github.com/kazmerdome/muzz/mocks"
	"github.com/stretchr/testify/assert"
)

type controllerFixture struct {
	controller explore_grpc.ExploreServiceServer
	mocks      struct {
		service *mocks.ExploreService
	}
	data struct {
		ctx context.Context
	}
}

func newControllerFixture(t *testing.T) *controllerFixture {
	f := &controllerFixture{}
	f.mocks.service = mocks.NewExploreService(t)
	f.data.ctx = context.TODO()

	f.controller = explore.NewExploreController(f.mocks.service)
	return f
}

// ListLikedYou

func TestCListLikedYou_FailsOn_service(t *testing.T) {
	f := newControllerFixture(t)
	now := fmt.Sprintf("%d", time.Now().Unix())
	req := &explore_grpc.ListLikedYouRequest{
		RecipientUserId: uuid.New().String(),
		PaginationToken: &now,
	}

	f.mocks.service.EXPECT().ListLikedYou(f.data.ctx, explore.ListLikedYouRequest{
		RecipientUserId: req.RecipientUserId,
		PaginationToken: req.PaginationToken,
	}).Return(nil, fmt.Errorf("an error"))

	c, err := f.controller.ListLikedYou(f.data.ctx, req)
	assert.EqualError(t, err, "an error")
	assert.Nil(t, c)
}
func TestCListLikedYou_SucceedWNil(t *testing.T) {
	f := newControllerFixture(t)
	now := fmt.Sprintf("%d", time.Now().Unix())
	req := &explore_grpc.ListLikedYouRequest{
		RecipientUserId: uuid.New().String(),
		PaginationToken: &now,
	}

	f.mocks.service.EXPECT().ListLikedYou(f.data.ctx, explore.ListLikedYouRequest{
		RecipientUserId: req.RecipientUserId,
		PaginationToken: req.PaginationToken,
	}).Return(nil, nil)

	c, err := f.controller.ListLikedYou(f.data.ctx, req)
	assert.NoError(t, err)
	assert.Nil(t, c)
}
func TestCListLikedYou_Succeed(t *testing.T) {
	f := newControllerFixture(t)
	now := fmt.Sprintf("%d", time.Now().Unix())
	req := &explore_grpc.ListLikedYouRequest{
		RecipientUserId: uuid.New().String(),
		PaginationToken: &now,
	}

	f.mocks.service.EXPECT().ListLikedYou(f.data.ctx, explore.ListLikedYouRequest{
		RecipientUserId: req.RecipientUserId,
		PaginationToken: req.PaginationToken,
	}).Return(&explore.ListLikedYouResponse{
		Likers: []explore.Liker{
			{
				ActorId:       uuid.NewString(),
				UnixTimestamp: time.Now().Unix(),
			},
		},
		NextPaginationToken: &now,
	}, nil)

	c, err := f.controller.ListLikedYou(f.data.ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, c)
}

// ListNewLikedYou

func TestCListNewLikedYou_FailsOn_service(t *testing.T) {
	f := newControllerFixture(t)
	now := fmt.Sprintf("%d", time.Now().Unix())
	req := &explore_grpc.ListLikedYouRequest{
		RecipientUserId: uuid.New().String(),
		PaginationToken: &now,
	}

	f.mocks.service.EXPECT().ListNewLikedYou(f.data.ctx, explore.ListLikedYouRequest{
		RecipientUserId: req.RecipientUserId,
		PaginationToken: req.PaginationToken,
	}).Return(nil, fmt.Errorf("an error"))

	c, err := f.controller.ListNewLikedYou(f.data.ctx, req)
	assert.EqualError(t, err, "an error")
	assert.Nil(t, c)
}
func TestCListNewLikedYou_SucceedWNil(t *testing.T) {
	f := newControllerFixture(t)
	now := fmt.Sprintf("%d", time.Now().Unix())
	req := &explore_grpc.ListLikedYouRequest{
		RecipientUserId: uuid.New().String(),
		PaginationToken: &now,
	}

	f.mocks.service.EXPECT().ListNewLikedYou(f.data.ctx, explore.ListLikedYouRequest{
		RecipientUserId: req.RecipientUserId,
		PaginationToken: req.PaginationToken,
	}).Return(nil, nil)

	c, err := f.controller.ListNewLikedYou(f.data.ctx, req)
	assert.NoError(t, err)
	assert.Nil(t, c)
}
func TestCListNewLikedYou_Succeed(t *testing.T) {
	f := newControllerFixture(t)
	now := fmt.Sprintf("%d", time.Now().Unix())
	req := &explore_grpc.ListLikedYouRequest{
		RecipientUserId: uuid.New().String(),
		PaginationToken: &now,
	}

	f.mocks.service.EXPECT().ListNewLikedYou(f.data.ctx, explore.ListLikedYouRequest{
		RecipientUserId: req.RecipientUserId,
		PaginationToken: req.PaginationToken,
	}).Return(&explore.ListLikedYouResponse{
		Likers: []explore.Liker{
			{
				ActorId:       uuid.NewString(),
				UnixTimestamp: time.Now().Unix(),
			},
		},
		NextPaginationToken: &now,
	}, nil)

	c, err := f.controller.ListNewLikedYou(f.data.ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, c)
}

// CountLikedYou

func TestCCountLikedYou_FailsOn_service(t *testing.T) {
	f := newControllerFixture(t)
	req := &explore_grpc.CountLikedYouRequest{
		RecipientUserId: uuid.New().String(),
	}
	f.mocks.service.EXPECT().CountLikedYou(f.data.ctx, explore.CountLikedYouRequest{
		RecipientUserId: req.RecipientUserId,
	}).Return(nil, fmt.Errorf("an error"))
	c, err := f.controller.CountLikedYou(f.data.ctx, req)
	assert.EqualError(t, err, "an error")
	assert.Nil(t, c)
}
func TestCCountLikedYou_Succeed(t *testing.T) {
	f := newControllerFixture(t)
	req := &explore_grpc.CountLikedYouRequest{
		RecipientUserId: uuid.New().String(),
	}
	f.mocks.service.EXPECT().CountLikedYou(f.data.ctx, explore.CountLikedYouRequest{
		RecipientUserId: req.RecipientUserId,
	}).Return(&explore.CountLikedYouResponse{
		Count: 42,
	}, nil)
	c, err := f.controller.CountLikedYou(f.data.ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, c.Count, uint64(42))
}

// PutDecision

func TestCPutDecision_FailsOn_service(t *testing.T) {
	f := newControllerFixture(t)
	req := &explore_grpc.PutDecisionRequest{
		ActorUserId:     uuid.New().String(),
		RecipientUserId: uuid.New().String(),
		LikedRecipient:  true,
	}
	f.mocks.service.EXPECT().PutDecision(f.data.ctx, explore.PutDecisionRequest{
		ActorUserId:     req.ActorUserId,
		RecipientUserId: req.RecipientUserId,
		LikedRecipient:  req.LikedRecipient,
	}).Return(nil, fmt.Errorf("an error"))
	c, err := f.controller.PutDecision(f.data.ctx, req)
	assert.EqualError(t, err, "an error")
	assert.Nil(t, c)
}
func TestCPutDecision_Succeed(t *testing.T) {
	f := newControllerFixture(t)
	req := &explore_grpc.PutDecisionRequest{
		ActorUserId:     uuid.New().String(),
		RecipientUserId: uuid.New().String(),
		LikedRecipient:  true,
	}
	f.mocks.service.EXPECT().PutDecision(f.data.ctx, explore.PutDecisionRequest{
		ActorUserId:     req.ActorUserId,
		RecipientUserId: req.RecipientUserId,
		LikedRecipient:  req.LikedRecipient,
	}).Return(&explore.PutDecisionResponse{
		MutualLikes: true,
	}, nil)
	c, err := f.controller.PutDecision(f.data.ctx, req)
	assert.NoError(t, err)
	assert.True(t, c.MutualLikes)
}
