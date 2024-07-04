package decision_test

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/kazmerdome/muzz/internal/module/decision"
	decisionQuerier "github.com/kazmerdome/muzz/internal/module/decision/decision-querier"
	"github.com/kazmerdome/muzz/mocks"
	"github.com/stretchr/testify/assert"
)

type repositoryFixture struct {
	repository decision.DecisionRepository
	mocks      struct {
		decisionQuerier *mocks.DecisionQuerier
	}
	data struct {
		ctx                    context.Context
		decision               decision.Decision
		upsertOneDto           decision.UpsertOneDto
		whereDto               decision.WhereDto
		querierDecision        decisionQuerier.Decision
		querierListParams      decisionQuerier.ListParams
		querierCountParams     decisionQuerier.CountParams
		querierUpsertOneParams decisionQuerier.UpsertOneParams
	}
}

func newRepositoryFixture(t *testing.T) *repositoryFixture {
	t.Parallel()
	f := &repositoryFixture{}
	f.mocks.decisionQuerier = mocks.NewDecisionQuerier(t)
	f.data.ctx = context.Background()
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
	timeNow := time.Now()
	pt := fmt.Sprintf("%d", timeNow.Unix())
	f.data.whereDto = decision.WhereDto{
		ActorUserID:         &f.data.decision.ActorUserID,
		RecipientUserID:     &f.data.decision.RecipientUserID,
		LikedRecipient:      &f.data.decision.LikedRecipient,
		RecipientLikesActor: &f.data.decision.RecipientLikesActor,
		PaginationToken:     &pt,
	}

	f.data.querierDecision = decisionQuerier.Decision{
		ID:                  f.data.decision.ID,
		ActorUserID:         f.data.decision.ActorUserID,
		RecipientUserID:     f.data.decision.RecipientUserID,
		LikedRecipient:      f.data.decision.LikedRecipient,
		RecipientLikesActor: f.data.decision.RecipientLikesActor,
		CreatedAt:           sql.NullTime{Valid: true, Time: f.data.decision.CreatedAt},
	}
	timestamp, _ := strconv.ParseInt(pt, 10, 64)
	f.data.querierListParams = decisionQuerier.ListParams{
		ActorUserID:         uuid.NullUUID{Valid: true, UUID: *f.data.whereDto.ActorUserID},
		RecipientUserID:     uuid.NullUUID{Valid: true, UUID: *f.data.whereDto.RecipientUserID},
		LikedRecipient:      sql.NullBool{Valid: true, Bool: *f.data.whereDto.LikedRecipient},
		RecipientLikesActor: sql.NullBool{Valid: true, Bool: *f.data.whereDto.RecipientLikesActor},
		PaginationToken:     sql.NullTime{Valid: true, Time: time.Unix(timestamp, 0)},
		Limit:               sql.NullInt32{Valid: true, Int32: 10},
	}
	f.data.querierUpsertOneParams = decisionQuerier.UpsertOneParams{
		ActorUserID:         f.data.decision.ActorUserID,
		RecipientUserID:     f.data.decision.RecipientUserID,
		LikedRecipient:      f.data.decision.LikedRecipient,
		RecipientLikesActor: f.data.decision.RecipientLikesActor,
	}
	f.data.querierCountParams = decisionQuerier.CountParams{
		ActorUserID:         f.data.querierListParams.ActorUserID,
		RecipientUserID:     f.data.querierListParams.RecipientUserID,
		LikedRecipient:      f.data.querierListParams.LikedRecipient,
		RecipientLikesActor: f.data.querierListParams.RecipientLikesActor,
	}
	f.repository = decision.NewDecisionRepository(f.mocks.decisionQuerier)
	return f
}

// UpsertOne
func TestUpsertOne_FailsOn_Querier(t *testing.T) {
	f := newRepositoryFixture(t)

	f.mocks.decisionQuerier.EXPECT().UpsertOne(f.data.ctx, f.data.querierUpsertOneParams).
		Return(f.data.querierDecision, fmt.Errorf("querier error"))

	c, err := f.repository.UpsertOne(f.data.ctx, f.data.upsertOneDto)
	assert.EqualError(t, err, "querier error")
	assert.Nil(t, c)
}

func TestUpsertOne_Succeed(t *testing.T) {
	f := newRepositoryFixture(t)

	f.mocks.decisionQuerier.EXPECT().UpsertOne(f.data.ctx, f.data.querierUpsertOneParams).
		Return(f.data.querierDecision, nil)

	c, err := f.repository.UpsertOne(f.data.ctx, f.data.upsertOneDto)
	assert.NoError(t, err)
	assert.NotNil(t, c)
}

// GetOneByActorUserId

func TestGetOneByActorUserId_FailsOn_querier(t *testing.T) {
	f := newRepositoryFixture(t)

	f.mocks.decisionQuerier.EXPECT().GetOneByActorUserId(f.data.ctx, f.data.querierDecision.ActorUserID).
		Return(f.data.querierDecision, fmt.Errorf("querier error"))

	c, err := f.repository.GetOneByActorUserId(f.data.ctx, f.data.decision.ActorUserID)
	assert.EqualError(t, err, "querier error")
	assert.Nil(t, c)
}

func TestGetOneByActorUserId_NotFound(t *testing.T) {
	f := newRepositoryFixture(t)

	f.mocks.decisionQuerier.EXPECT().GetOneByActorUserId(f.data.ctx, f.data.querierDecision.RecipientUserID).
		Return(f.data.querierDecision, sql.ErrNoRows)

	c, err := f.repository.GetOneByActorUserId(f.data.ctx, f.data.decision.RecipientUserID)
	assert.EqualError(t, err, decision.ErrNotFound.Error())
	assert.Nil(t, c)
}

func TestGetOneByActorUserId_Succeed(t *testing.T) {
	f := newRepositoryFixture(t)

	f.mocks.decisionQuerier.EXPECT().GetOneByActorUserId(f.data.ctx, f.data.querierDecision.ActorUserID).
		Return(f.data.querierDecision, nil)

	c, err := f.repository.GetOneByActorUserId(f.data.ctx, f.data.decision.ActorUserID)
	assert.NoError(t, err)
	assert.NotNil(t, c)
}

// GetOneByRecipientUserID

func TestGetOneByRecipientUserID_FailsOn_querier(t *testing.T) {
	f := newRepositoryFixture(t)

	f.mocks.decisionQuerier.EXPECT().GetOneByRecipientUserID(f.data.ctx, f.data.querierDecision.RecipientUserID).
		Return(f.data.querierDecision, fmt.Errorf("querier error"))

	c, err := f.repository.GetOneByRecipientUserID(f.data.ctx, f.data.decision.RecipientUserID)
	assert.EqualError(t, err, "querier error")
	assert.Nil(t, c)
}

func TestGetOneByRecipientUserID_NotFound(t *testing.T) {
	f := newRepositoryFixture(t)

	f.mocks.decisionQuerier.EXPECT().GetOneByRecipientUserID(f.data.ctx, f.data.querierDecision.RecipientUserID).
		Return(f.data.querierDecision, sql.ErrNoRows)

	c, err := f.repository.GetOneByRecipientUserID(f.data.ctx, f.data.decision.RecipientUserID)
	assert.EqualError(t, err, decision.ErrNotFound.Error())
	assert.Nil(t, c)
}

func TestGetOneByRecipientUserID_Succeed(t *testing.T) {
	f := newRepositoryFixture(t)

	f.mocks.decisionQuerier.EXPECT().GetOneByRecipientUserID(f.data.ctx, f.data.querierDecision.RecipientUserID).
		Return(f.data.querierDecision, nil)

	c, err := f.repository.GetOneByRecipientUserID(f.data.ctx, f.data.decision.RecipientUserID)
	assert.NoError(t, err)
	assert.NotNil(t, c)
}

// List

func TestList_FailsOn_Querier(t *testing.T) {
	f := newRepositoryFixture(t)

	f.mocks.decisionQuerier.EXPECT().List(f.data.ctx, f.data.querierListParams).
		Return([]decisionQuerier.Decision{}, fmt.Errorf("querier error"))

	c, err := f.repository.List(f.data.ctx, &f.data.whereDto)
	assert.EqualError(t, err, "querier error")
	assert.Nil(t, c)
}

func TestList_FailsOn_NotValidTimeStamp(t *testing.T) {
	f := newRepositoryFixture(t)
	nt := "not-a-valid-unix"
	f.data.whereDto.PaginationToken = &nt
	c, err := f.repository.List(f.data.ctx, &f.data.whereDto)
	assert.EqualError(t, err, decision.ErrPaginationTokenIsNotUnixTimeStamp.Error())
	assert.Nil(t, c)
}

func TestList_Succeed(t *testing.T) {
	f := newRepositoryFixture(t)

	f.mocks.decisionQuerier.EXPECT().List(f.data.ctx, f.data.querierListParams).
		Return([]decisionQuerier.Decision{f.data.querierDecision}, nil)

	c, err := f.repository.List(f.data.ctx, &f.data.whereDto)
	assert.NoError(t, err)
	assert.NotEmpty(t, c)
}

// Count

func TestCount_FailsOn_Querier(t *testing.T) {
	f := newRepositoryFixture(t)
	f.mocks.decisionQuerier.EXPECT().Count(f.data.ctx, f.data.querierCountParams).
		Return(0, fmt.Errorf("querier error"))
	c, err := f.repository.Count(f.data.ctx, &f.data.whereDto)
	assert.EqualError(t, err, "querier error")
	assert.Zero(t, c)
}

func TestCount_Succeed(t *testing.T) {
	f := newRepositoryFixture(t)
	f.mocks.decisionQuerier.EXPECT().Count(f.data.ctx, f.data.querierCountParams).
		Return(42, nil)
	c, err := f.repository.Count(f.data.ctx, &f.data.whereDto)
	assert.NoError(t, err)
	assert.Equal(t, c, int64(42))
}
