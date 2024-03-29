package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.29

import (
	"context"

	"github.com/AI1411/go-gql-server/graph/model"
	"github.com/AI1411/go-gql-server/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

// CreateTweet is the resolver for the createTweet field.
func (r *mutationResolver) CreateTweet(ctx context.Context, input model.CreateTweetInput) (string, error) {
	res, err := r.TweetClient.CreateTweet(ctx, &grpc.CreateTweetRequest{
		UserId: input.UserID,
		Body:   input.Body,
	})
	if err != nil {
		return "", err
	}

	return res.GetId(), nil
}

// ListTweet is the resolver for the listTweet field.
func (r *queryResolver) ListTweet(ctx context.Context, input *string) ([]*model.Tweet, error) {
	res, err := r.TweetClient.ListTweet(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	tweets := make([]*model.Tweet, len(res.Tweets))
	for i, t := range res.Tweets {
		tweets[i] = &model.Tweet{
			ID:        t.Id,
			Body:      t.Body,
			UserID:    t.UserId,
			CreatedAt: t.CreatedAt.AsTime().String(),
			UpdatedAt: t.UpdatedAt.AsTime().String(),
		}
	}

	return tweets, nil
}
