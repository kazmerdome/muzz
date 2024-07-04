package explore

import "context"

// Provider

type ExploreService interface {
	ListLikedYou(ctx context.Context, req ListLikedYouRequest) (*ListLikedYouResponse, error)    // List all users who liked the recipient
	ListNewLikedYou(ctx context.Context, req ListLikedYouRequest) (*ListLikedYouResponse, error) // List all users who liked the recipient excluding those who have been liked in return
	CountLikedYou(ctx context.Context, req CountLikedYouRequest) (*CountLikedYouResponse, error) // Count the number of users who liked the recipient
	PutDecision(ctx context.Context, req PutDecisionRequest) (*PutDecisionResponse, error)       // Record the decision of the actor to like or pass the recipient
}

// Entity & DTO

type Liker struct {
	ActorId       string `json:"actor_id"`
	UnixTimestamp int64  `json:"unix_timestamp"`
}

type ListLikedYouRequest struct {
	RecipientUserId string  `json:"recipient_user_id"`
	PaginationToken *string `json:"pagination_token"`
}

type ListLikedYouResponse struct {
	Likers              []Liker `json:"likers"`
	NextPaginationToken *string `json:"next_pagination_token"`
}

type CountLikedYouRequest struct {
	RecipientUserId string `json:"recipient_user_id"`
}

type CountLikedYouResponse struct {
	Count int64 `json:"count"`
}

type PutDecisionRequest struct {
	ActorUserId     string `json:"actor_user_id"`
	RecipientUserId string `json:"recipient_user_id"`
	LikedRecipient  bool   `json:"liked_recipient"`
}

type PutDecisionResponse struct {
	MutualLikes bool `json:"mutual_likes"`
}
