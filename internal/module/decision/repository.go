package decision

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"time"

	"github.com/google/uuid"
	decisionQuerier "github.com/kazmerdome/muzz/internal/module/decision/decision-querier"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type decisionRepository struct {
	querier      decisionQuerier.Querier
	logger       zerolog.Logger
	defaultLimit int32
}

func NewDecisionRepository(querier decisionQuerier.Querier) *decisionRepository {
	return &decisionRepository{
		querier:      querier,
		defaultLimit: 10,
		logger: log.
			With().
			Str("module", "decision").
			Str("provider", "repository").
			Logger(),
	}
}

func (r *decisionRepository) UpsertOne(ctx context.Context, dto UpsertOneDto) (*Decision, error) {
	params := decisionQuerier.UpsertOneParams{
		ActorUserID:         dto.ActorUserID,
		RecipientUserID:     dto.RecipientUserID,
		LikedRecipient:      dto.LikedRecipient,
		RecipientLikesActor: dto.RecipientLikesActor,
	}

	qd, err := r.querier.UpsertOne(ctx, params)
	if err != nil {
		r.logger.
			Error().
			Err(err).
			Str("method", "UpsertOne").
			Str("event", "call querier.UpsertOne").
			Send()
		return nil, err
	}

	return r.buildModelFromQuerier(qd), nil
}

func (r *decisionRepository) GetOneByActorUserId(ctx context.Context, actorUserID uuid.UUID) (*Decision, error) {
	model, err := r.querier.GetOneByActorUserId(ctx, actorUserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		r.logger.
			Error().
			Err(err).
			Str("method", "GetOneByRecipientUserID").
			Str("event", "call querier.GetOneByRecipientUserID").
			Send()
		return nil, err
	}
	return r.buildModelFromQuerier(model), nil
}

func (r *decisionRepository) GetOneByRecipientUserID(ctx context.Context, recipientUserID uuid.UUID) (*Decision, error) {
	model, err := r.querier.GetOneByRecipientUserID(ctx, recipientUserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		r.logger.
			Error().
			Err(err).
			Str("method", "GetOneByRecipientUserID").
			Str("event", "call querier.GetOneByRecipientUserID").
			Send()
		return nil, err
	}
	return r.buildModelFromQuerier(model), nil
}

func (r *decisionRepository) List(ctx context.Context, where *WhereDto) ([]Decision, error) {
	querierParams := decisionQuerier.ListParams{
		Limit: sql.NullInt32{Valid: true, Int32: r.defaultLimit},
	}
	if where != nil {
		if where.ActorUserID != nil {
			querierParams.ActorUserID = uuid.NullUUID{Valid: true, UUID: *where.ActorUserID}
		}
		if where.RecipientUserID != nil {
			querierParams.RecipientUserID = uuid.NullUUID{Valid: true, UUID: *where.RecipientUserID}
		}
		if where.LikedRecipient != nil {
			querierParams.LikedRecipient = sql.NullBool{Valid: true, Bool: *where.LikedRecipient}
		}
		if where.RecipientLikesActor != nil {
			querierParams.RecipientLikesActor = sql.NullBool{Valid: true, Bool: *where.RecipientLikesActor}
		}
		if where.PaginationToken != nil {
			isValid, timestamp := r.isUnixTimestamp(*where.PaginationToken)
			if !isValid {
				return nil, ErrPaginationTokenIsNotUnixTimeStamp
			}
			querierParams.PaginationToken = sql.NullTime{Valid: true, Time: time.Unix(timestamp, 0)}
		}
	}
	models, err := r.querier.List(ctx, querierParams)
	if err != nil {
		r.logger.
			Error().
			Err(err).
			Str("method", "List").
			Str("category", "call querier.List").
			Send()
		return nil, err
	}
	items := make([]Decision, len(models))
	for i, model := range models {
		qd := r.buildModelFromQuerier(model)
		items[i] = *qd
	}
	return items, nil
}

func (r *decisionRepository) Count(ctx context.Context, where *WhereDto) (int64, error) {
	querierParams := decisionQuerier.CountParams{}
	if where != nil {
		if where.ActorUserID != nil {
			querierParams.ActorUserID = uuid.NullUUID{Valid: true, UUID: *where.ActorUserID}
		}
		if where.RecipientUserID != nil {
			querierParams.RecipientUserID = uuid.NullUUID{Valid: true, UUID: *where.RecipientUserID}
		}
		if where.LikedRecipient != nil {
			querierParams.LikedRecipient = sql.NullBool{Valid: true, Bool: *where.LikedRecipient}
		}
		if where.RecipientLikesActor != nil {
			querierParams.RecipientLikesActor = sql.NullBool{Valid: true, Bool: *where.RecipientLikesActor}
		}
	}
	count, err := r.querier.Count(ctx, querierParams)
	if err != nil {
		r.logger.
			Error().
			Err(err).
			Str("method", "Count").
			Str("category", "call querier.Count").
			Send()
		return 0, err
	}
	return count, nil
}

// Helpers
//

func (r *decisionRepository) buildModelFromQuerier(qd decisionQuerier.Decision) *Decision {
	return &Decision{
		ID:              qd.ID,
		ActorUserID:     qd.ActorUserID,
		RecipientUserID: qd.RecipientUserID,
		LikedRecipient:  qd.LikedRecipient,
		CreatedAt:       qd.CreatedAt.Time,
	}
}

func (r *decisionRepository) isUnixTimestamp(str string) (bool, int64) {
	// Attempt to parse the string as an integer
	timestamp, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return false, timestamp // Not a valid integer
	}

	// Validate if the parsed integer is a Unix timestamp
	// In Unix, the minimum timestamp is January 1, 1970 UTC
	// You can adjust this if you expect negative timestamps or timestamps before 1970
	minUnixTime := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()
	return timestamp >= minUnixTime, timestamp
}
