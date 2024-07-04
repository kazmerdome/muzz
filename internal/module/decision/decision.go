package decision

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Error

var ErrPaginationTokenIsNotUnixTimeStamp = fmt.Errorf("invalid pagination token")
var ErrNotFound = fmt.Errorf("decision is not found")

// Provider

type DecisionRepository interface {
	UpsertOne(ctx context.Context, dto UpsertOneDto) (*Decision, error)
	GetOneByActorUserId(ctx context.Context, actorUserID uuid.UUID) (*Decision, error)
	GetOneByRecipientUserID(ctx context.Context, recipientUserID uuid.UUID) (*Decision, error)
	List(ctx context.Context, where *WhereDto) ([]Decision, error)
	Count(ctx context.Context, where *WhereDto) (int64, error)
}

// Entity

type Decision struct {
	ID                  uuid.UUID `json:"id"`
	ActorUserID         uuid.UUID `json:"actor_user_id"`
	RecipientUserID     uuid.UUID `json:"recipient_user_id"`
	LikedRecipient      bool      `json:"liked_recipient"`
	RecipientLikesActor bool      `json:"recipient_likes_actor"`
	CreatedAt           time.Time `json:"createdAt"`
}

// DTO

type UpsertOneDto struct {
	ActorUserID         uuid.UUID `json:"actor_user_id"`
	RecipientUserID     uuid.UUID `json:"recipient_user_id"`
	LikedRecipient      bool      `json:"liked_recipient"`
	RecipientLikesActor bool      `json:"recipient_likes_actor"`
}

type WhereDto struct {
	ActorUserID         *uuid.UUID `json:"actor_user_id"`
	RecipientUserID     *uuid.UUID `json:"recipient_user_id"`
	LikedRecipient      *bool      `json:"liked_recipient"`
	RecipientLikesActor *bool      `json:"recipient_likes_actor"`
	PaginationToken     *string    `json:"pagination_token"`
}
