package explore

import (
	context "context"

	explore_grpc "github.com/kazmerdome/muzz/internal/module/explore/explore-grpc"
)

// implements explore_grpc.ExploreServiceServer

type exploreController struct {
	explore_grpc.UnimplementedExploreServiceServer
	service ExploreService
}

func NewExploreController(service ExploreService) *exploreController {
	return &exploreController{
		service: service,
	}
}

func (r *exploreController) ListLikedYou(ctx context.Context, req *explore_grpc.ListLikedYouRequest) (*explore_grpc.ListLikedYouResponse, error) {
	listRes, err := r.service.ListLikedYou(ctx, ListLikedYouRequest{
		RecipientUserId: req.RecipientUserId,
		PaginationToken: req.PaginationToken,
	})
	if err != nil {
		return nil, err
	}
	if listRes == nil {
		return nil, nil
	}
	resp := &explore_grpc.ListLikedYouResponse{
		NextPaginationToken: listRes.NextPaginationToken,
		Likers:              []*explore_grpc.ListLikedYouResponse_Liker{},
	}
	for _, l := range listRes.Likers {
		resp.Likers = append(resp.Likers, &explore_grpc.ListLikedYouResponse_Liker{
			ActorId:       l.ActorId,
			UnixTimestamp: uint64(l.UnixTimestamp),
		})
	}
	return resp, nil
}

func (r *exploreController) ListNewLikedYou(ctx context.Context, req *explore_grpc.ListLikedYouRequest) (*explore_grpc.ListLikedYouResponse, error) {
	listRes, err := r.service.ListNewLikedYou(ctx, ListLikedYouRequest{
		RecipientUserId: req.RecipientUserId,
		PaginationToken: req.PaginationToken,
	})
	if err != nil {
		return nil, err
	}
	if listRes == nil {
		return nil, nil
	}
	resp := &explore_grpc.ListLikedYouResponse{
		NextPaginationToken: listRes.NextPaginationToken,
		Likers:              []*explore_grpc.ListLikedYouResponse_Liker{},
	}
	for _, l := range listRes.Likers {
		resp.Likers = append(resp.Likers, &explore_grpc.ListLikedYouResponse_Liker{
			ActorId:       l.ActorId,
			UnixTimestamp: uint64(l.UnixTimestamp),
		})
	}
	return resp, nil
}

func (r *exploreController) CountLikedYou(ctx context.Context, req *explore_grpc.CountLikedYouRequest) (*explore_grpc.CountLikedYouResponse, error) {
	c, err := r.service.CountLikedYou(ctx, CountLikedYouRequest{
		RecipientUserId: req.RecipientUserId,
	})
	if err != nil {
		return nil, err
	}
	return &explore_grpc.CountLikedYouResponse{
		Count: uint64(c.Count),
	}, nil
}

func (r *exploreController) PutDecision(ctx context.Context, req *explore_grpc.PutDecisionRequest) (*explore_grpc.PutDecisionResponse, error) {
	sr, err := r.service.PutDecision(ctx, PutDecisionRequest{
		ActorUserId:     req.ActorUserId,
		RecipientUserId: req.RecipientUserId,
		LikedRecipient:  req.LikedRecipient,
	})
	if err != nil {
		return nil, err
	}

	return &explore_grpc.PutDecisionResponse{
		MutualLikes: sr.MutualLikes,
	}, nil
}
