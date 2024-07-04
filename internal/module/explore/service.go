package explore

import (
	context "context"
	"fmt"

	"github.com/google/uuid"
	"github.com/kazmerdome/muzz/internal/module/decision"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type exploreService struct {
	logger             zerolog.Logger
	decisionRepository decision.DecisionRepository
}

func NewExploreService(decisionRepository decision.DecisionRepository) *exploreService {
	return &exploreService{
		decisionRepository: decisionRepository,
		logger: log.
			With().
			Str("module", "explore").
			Str("provider", "service").
			Logger(),
	}
}

// List all users who liked the recipient
func (r *exploreService) ListLikedYou(ctx context.Context, req ListLikedYouRequest) (*ListLikedYouResponse, error) {
	// parse req.RecipientUserId to uuid.UUID
	uid, err := uuid.Parse(req.RecipientUserId)
	if err != nil {
		r.logger.
			Error().
			Err(err).
			Str("method", "ListLikedYou").
			Str("event", "parse req.RecipientUserId to uuid").
			Send()
		return nil, err
	}
	liked := true

	// Call decision repository
	decisions, err := r.decisionRepository.List(ctx, &decision.WhereDto{
		PaginationToken: req.PaginationToken,
		RecipientUserID: &uid,
		LikedRecipient:  &liked,
	})
	if err != nil {
		r.logger.
			Error().
			Err(err).
			Str("method", "ListLikedYou").
			Str("event", "decisionRepository.List").
			Send()
		return nil, err
	}

	// Prepare response
	resp := &ListLikedYouResponse{}
	if len(decisions) > 0 {
		token := fmt.Sprintf("%d", (decisions[len(decisions)-1].CreatedAt.Unix()))
		resp.NextPaginationToken = &token
		for _, d := range decisions {
			resp.Likers = append(resp.Likers, Liker{
				ActorId:       d.ActorUserID.String(),
				UnixTimestamp: d.CreatedAt.Unix(),
			})
		}
	}
	return resp, nil
}

// List all users who liked the recipient excluding those who have been liked in return
func (r *exploreService) ListNewLikedYou(ctx context.Context, req ListLikedYouRequest) (*ListLikedYouResponse, error) {
	// parse req.RecipientUserId to uuid.UUID
	uid, err := uuid.Parse(req.RecipientUserId)
	if err != nil {
		r.logger.
			Error().
			Err(err).
			Str("method", "ListNewLikedYou").
			Str("event", "parse req.RecipientUserId to uuid").
			Send()
		return nil, err
	}

	// Call decision repository
	liked := true
	nope := false
	decisions, err := r.decisionRepository.List(ctx, &decision.WhereDto{
		PaginationToken:     req.PaginationToken,
		RecipientUserID:     &uid,
		LikedRecipient:      &liked,
		RecipientLikesActor: &nope,
	})
	if err != nil {
		r.logger.
			Error().
			Err(err).
			Str("method", "ListNewLikedYou").
			Str("event", "decisionRepository.List").
			Send()
		return nil, err
	}

	// Prepare response
	resp := &ListLikedYouResponse{}
	if len(decisions) > 0 {
		token := fmt.Sprintf("%d", (decisions[len(decisions)-1].CreatedAt.Unix()))
		resp.NextPaginationToken = &token
		for _, d := range decisions {
			resp.Likers = append(resp.Likers, Liker{
				ActorId:       d.ActorUserID.String(),
				UnixTimestamp: d.CreatedAt.Unix(),
			})
		}
	}
	return resp, nil
}

// Count the number of users who liked the recipient
func (r *exploreService) CountLikedYou(ctx context.Context, req CountLikedYouRequest) (*CountLikedYouResponse, error) {
	// parse uuids
	ruid, err := uuid.Parse(req.RecipientUserId)
	if err != nil {
		r.logger.
			Error().
			Err(err).
			Str("method", "CountLikedYou").
			Str("event", "parse req.RecipientUserId to uuid").
			Send()
		return nil, err
	}

	// Call decision repository
	liked := true
	c, err := r.decisionRepository.Count(ctx, &decision.WhereDto{
		RecipientUserID: &ruid,
		LikedRecipient:  &liked,
	})
	if err != nil {
		r.logger.
			Error().
			Err(err).
			Str("method", "CountLikedYou").
			Str("event", "call decisionRepository.Count").
			Send()
		return nil, err
	}

	return &CountLikedYouResponse{
		Count: c,
	}, nil
}

func (r *exploreService) PutDecision(ctx context.Context, req PutDecisionRequest) (*PutDecisionResponse, error) {
	// parse uuids
	auid, err := uuid.Parse(req.ActorUserId)
	if err != nil {
		r.logger.
			Error().
			Err(err).
			Str("method", "PutDecision").
			Str("event", "parse req.ActorUserId to uuid").
			Send()
		return nil, err
	}
	ruid, err := uuid.Parse(req.RecipientUserId)
	if err != nil {
		r.logger.
			Error().
			Err(err).
			Str("method", "PutDecision").
			Str("event", "parse req.RecipientUserId to uuid").
			Send()
		return nil, err
	}

	// Check if Recipient is already liked the Actor
	matches, err := r.decisionRepository.List(ctx, &decision.WhereDto{
		ActorUserID:     &ruid,
		RecipientUserID: &auid,
	})
	if err != nil {
		r.logger.
			Error().
			Err(err).
			Str("method", "PutDecision").
			Str("event", "call decisionRepository.Count").
			Send()
		return nil, err
	}
	dto := decision.UpsertOneDto{
		ActorUserID:     auid,
		RecipientUserID: ruid,
		LikedRecipient:  req.LikedRecipient,
	}

	if len(matches) > 0 {
		// I assume that it is only one item in the matches
		match := matches[0]

		// if we have a match and Actor likes the Recipient -> set RecipientLikesActor to true
		if match.LikedRecipient {
			dto.RecipientLikesActor = true
		}

		// if we have a match and Actor does not like the Recipient -> update the match as well
		_, err := r.decisionRepository.UpsertOne(ctx, decision.UpsertOneDto{
			ActorUserID:         match.ActorUserID,
			RecipientUserID:     match.RecipientUserID,
			LikedRecipient:      match.LikedRecipient,
			RecipientLikesActor: req.LikedRecipient,
		})
		if err != nil {
			r.logger.
				Error().
				Err(err).
				Str("method", "PutDecision").
				Str("event", "update recipient already existing decision decisionRepository.UpsertOne").
				Send()
			return nil, err
		}

	}

	// Call the upsert repository method
	dec, err := r.decisionRepository.UpsertOne(ctx, dto)
	if err != nil {
		r.logger.
			Error().
			Err(err).
			Str("method", "PutDecision").
			Str("event", "call decisionRepository.UpsertOne").
			Send()
		return nil, err
	}

	// Check for mutual likes
	resp := &PutDecisionResponse{}
	if dec.RecipientLikesActor {
		resp.MutualLikes = true
	}
	return resp, nil
}
