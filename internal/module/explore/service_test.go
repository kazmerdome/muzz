package explore_test

import (
	context "context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/kazmerdome/muzz/internal/module/decision"
	"github.com/kazmerdome/muzz/internal/module/explore"
	"github.com/kazmerdome/muzz/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type serviceFixture struct {
	service explore.ExploreService
	mocks   struct {
		decisionRepository *mocks.DecisionRepository
	}
	data struct {
		ctx                   context.Context
		decision              decision.Decision
		upsertOneDto          decision.UpsertOneDto
		whereDto              decision.WhereDto
		listLikedYouRequest   explore.ListLikedYouRequest
		listLikedYouResponse  explore.ListLikedYouResponse
		countLikedYouRequest  explore.CountLikedYouRequest
		countLikedYouResponse explore.CountLikedYouResponse
		putDecisionRequest    explore.PutDecisionRequest
		putDecisionResponse   explore.PutDecisionResponse
	}
}

func newServiceFixture(t *testing.T) *serviceFixture {
	f := &serviceFixture{}
	f.mocks.decisionRepository = mocks.NewDecisionRepository(t)
	f.data.ctx = context.TODO()
	pt := time.Now().Unix()
	spt := fmt.Sprintf("%d", pt)

	f.data.decision = decision.Decision{
		ID:                  uuid.New(),
		ActorUserID:         uuid.New(),
		RecipientUserID:     uuid.New(),
		LikedRecipient:      true,
		RecipientLikesActor: false,
		CreatedAt:           time.Now(),
	}
	f.data.upsertOneDto = decision.UpsertOneDto{
		ActorUserID:         f.data.decision.ActorUserID,
		RecipientUserID:     f.data.decision.RecipientUserID,
		LikedRecipient:      f.data.decision.LikedRecipient,
		RecipientLikesActor: f.data.decision.RecipientLikesActor,
	}
	f.data.whereDto = decision.WhereDto{
		ActorUserID:         &f.data.decision.ActorUserID,
		RecipientUserID:     &f.data.decision.RecipientUserID,
		LikedRecipient:      &f.data.decision.LikedRecipient,
		RecipientLikesActor: &f.data.decision.RecipientLikesActor,
		PaginationToken:     &spt,
	}
	f.data.listLikedYouRequest = explore.ListLikedYouRequest{
		RecipientUserId: f.data.decision.RecipientUserID.String(),
		PaginationToken: &spt,
	}
	f.data.listLikedYouResponse = explore.ListLikedYouResponse{
		Likers: []explore.Liker{
			{
				ActorId:       f.data.decision.ActorUserID.String(),
				UnixTimestamp: pt,
			},
		},
		NextPaginationToken: &spt,
	}

	f.data.countLikedYouRequest = explore.CountLikedYouRequest{
		RecipientUserId: f.data.listLikedYouRequest.RecipientUserId,
	}
	f.data.countLikedYouResponse = explore.CountLikedYouResponse{
		Count: 55,
	}
	f.data.putDecisionRequest = explore.PutDecisionRequest{
		ActorUserId:     f.data.listLikedYouResponse.Likers[0].ActorId,
		RecipientUserId: f.data.listLikedYouRequest.RecipientUserId,
		LikedRecipient:  true,
	}
	f.data.putDecisionResponse = explore.PutDecisionResponse{
		MutualLikes: false,
	}

	f.service = explore.NewExploreService(f.mocks.decisionRepository)
	return f
}

// ListLikedYou

func TestListLikedYou_FailsOn_uuidParse(t *testing.T) {
	f := newServiceFixture(t)
	f.data.listLikedYouRequest.RecipientUserId = "not-valid"
	c, err := f.service.ListLikedYou(f.data.ctx, f.data.listLikedYouRequest)
	assert.EqualError(t, err, "invalid UUID length: 9")
	assert.Nil(t, c)
}

func TestListLikedYou_FailsOn_decisionRepository(t *testing.T) {
	f := newServiceFixture(t)

	f.mocks.decisionRepository.EXPECT().List(f.data.ctx, mock.Anything).Return(nil, fmt.Errorf("error"))

	c, err := f.service.ListLikedYou(f.data.ctx, f.data.listLikedYouRequest)
	assert.EqualError(t, err, "error")
	assert.Nil(t, c)
}

func TestListLikedYou_Succeed(t *testing.T) {
	f := newServiceFixture(t)

	f.mocks.decisionRepository.EXPECT().List(f.data.ctx, mock.Anything).Return([]decision.Decision{f.data.decision}, nil)

	c, err := f.service.ListLikedYou(f.data.ctx, f.data.listLikedYouRequest)
	assert.NoError(t, err, "error")
	assert.NotEmpty(t, c)
}

// ListNewLikedYou

func TestListNewLikedYou_FailsOn_uuidParse(t *testing.T) {
	f := newServiceFixture(t)
	f.data.listLikedYouRequest.RecipientUserId = "not-valid"
	c, err := f.service.ListNewLikedYou(f.data.ctx, f.data.listLikedYouRequest)
	assert.EqualError(t, err, "invalid UUID length: 9")
	assert.Nil(t, c)
}

func TestListNewLikedYou_FailsOn_decisionRepository(t *testing.T) {
	f := newServiceFixture(t)

	f.mocks.decisionRepository.EXPECT().List(f.data.ctx, mock.Anything).Return(nil, fmt.Errorf("error"))

	c, err := f.service.ListNewLikedYou(f.data.ctx, f.data.listLikedYouRequest)
	assert.EqualError(t, err, "error")
	assert.Nil(t, c)
}

func TestListNewLikedYou_Succeed(t *testing.T) {
	f := newServiceFixture(t)

	f.mocks.decisionRepository.EXPECT().List(f.data.ctx, mock.Anything).Return([]decision.Decision{f.data.decision}, nil)

	c, err := f.service.ListNewLikedYou(f.data.ctx, f.data.listLikedYouRequest)
	assert.NoError(t, err, "error")
	assert.NotEmpty(t, c)
}

// CountLikedYou

func TestCountLikedYou_FailsOn_uuidParse(t *testing.T) {
	f := newServiceFixture(t)
	f.data.countLikedYouRequest.RecipientUserId = "not-valid"
	c, err := f.service.CountLikedYou(f.data.ctx, f.data.countLikedYouRequest)
	assert.EqualError(t, err, "invalid UUID length: 9")
	assert.Nil(t, c)
}

func TestCountLikedYou_FailsOn_decisionRepository(t *testing.T) {
	f := newServiceFixture(t)

	f.mocks.decisionRepository.EXPECT().Count(f.data.ctx, mock.Anything).Return(int64(42), fmt.Errorf("error"))

	c, err := f.service.CountLikedYou(f.data.ctx, f.data.countLikedYouRequest)
	assert.EqualError(t, err, "error")
	assert.Zero(t, c)
}
func TestCountLikedYou_Succeed(t *testing.T) {
	f := newServiceFixture(t)

	f.mocks.decisionRepository.EXPECT().Count(f.data.ctx, mock.Anything).Return(int64(42), nil)

	c, err := f.service.CountLikedYou(f.data.ctx, f.data.countLikedYouRequest)
	assert.NoError(t, err, "error")
	assert.Equal(t, c.Count, int64(42))
}

// PutDecision

func TestPutDecision_FailsOn_uuidParse(t *testing.T) {
	f := newServiceFixture(t)
	f.data.putDecisionRequest.RecipientUserId = "not-valid"
	c, err := f.service.PutDecision(f.data.ctx, f.data.putDecisionRequest)
	assert.EqualError(t, err, "invalid UUID length: 9")
	assert.Nil(t, c)
}

func TestPutDecision_FailsOn_uuidParse2(t *testing.T) {
	f := newServiceFixture(t)
	f.data.putDecisionRequest.ActorUserId = "not-valid"
	c, err := f.service.PutDecision(f.data.ctx, f.data.putDecisionRequest)
	assert.EqualError(t, err, "invalid UUID length: 9")
	assert.Nil(t, c)
}

func TestPutDecision_FailsOn_List(t *testing.T) {
	f := newServiceFixture(t)
	f.mocks.decisionRepository.EXPECT().List(f.data.ctx, mock.Anything).Return(nil, fmt.Errorf("error"))

	c, err := f.service.PutDecision(f.data.ctx, f.data.putDecisionRequest)
	assert.EqualError(t, err, "error")
	assert.Nil(t, c)
}

func TestPutDecision_FailsOn_UpsertOne(t *testing.T) {
	f := newServiceFixture(t)
	f.mocks.decisionRepository.EXPECT().List(f.data.ctx, mock.Anything).Return([]decision.Decision{f.data.decision}, nil)
	f.mocks.decisionRepository.EXPECT().UpsertOne(f.data.ctx, mock.Anything).Return(nil, fmt.Errorf("upsert error"))

	c, err := f.service.PutDecision(f.data.ctx, f.data.putDecisionRequest)
	assert.EqualError(t, err, "upsert error")
	assert.Nil(t, c)
}

func TestPutDecision_FailsOn_UpsertOne2(t *testing.T) {
	f := newServiceFixture(t)
	f.mocks.decisionRepository.EXPECT().List(f.data.ctx, mock.Anything).Return([]decision.Decision{}, nil)
	f.mocks.decisionRepository.EXPECT().UpsertOne(f.data.ctx, decision.UpsertOneDto{
		ActorUserID:     f.data.decision.ActorUserID,
		RecipientUserID: f.data.decision.RecipientUserID,
		LikedRecipient:  f.data.decision.LikedRecipient,
	}).Return(nil, fmt.Errorf("upsert error"))

	c, err := f.service.PutDecision(f.data.ctx, f.data.putDecisionRequest)
	assert.EqualError(t, err, "upsert error")
	assert.Nil(t, c)
}

func TestPutDecision_Succeed(t *testing.T) {
	f := newServiceFixture(t)
	f.mocks.decisionRepository.EXPECT().List(f.data.ctx, mock.Anything).Return([]decision.Decision{}, nil)
	f.mocks.decisionRepository.EXPECT().UpsertOne(f.data.ctx, decision.UpsertOneDto{
		ActorUserID:     f.data.decision.ActorUserID,
		RecipientUserID: f.data.decision.RecipientUserID,
		LikedRecipient:  f.data.decision.LikedRecipient,
	}).Return(&decision.Decision{RecipientLikesActor: true}, nil)

	c, err := f.service.PutDecision(f.data.ctx, f.data.putDecisionRequest)
	assert.NoError(t, err)
	assert.True(t, c.MutualLikes)
}
