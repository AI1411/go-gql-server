package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.26

import (
	"context"
	"log"

	"github.com/AI1411/go-gql-server/grpc"

	"github.com/AI1411/go-gql-server/graph/model"
)

// CreateRoom is the resolver for the createRoom field.
func (r *mutationResolver) CreateRoom(ctx context.Context, input model.CreateRoomInput) (string, error) {
	res, err := r.RoomClient.CreateRoom(ctx, &grpc.CreateRoomRequest{
		UserId: input.UserID,
	})
	if err != nil {
		return "", err
	}

	return res.GetId(), nil
}

// DeleteRoom is the resolver for the deleteRoom field.
func (r *mutationResolver) DeleteRoom(ctx context.Context, input model.DeleteRoomInput) (*string, error) {
	_, err := r.RoomClient.DeleteRoom(ctx, &grpc.DeleteRoomRequest{
		Id: input.ID,
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// ListRoom is the resolver for the listRoom field.
func (r *queryResolver) ListRoom(ctx context.Context, input model.ListRoomInput) ([]*model.Room, error) {
	res, err := r.RoomClient.ListRoom(ctx, &grpc.ListRoomRequest{
		UserId: input.UserID,
	})

	log.Printf("res: %+v", res)

	if err != nil {
		return nil, err
	}

	rooms := make([]*model.Room, len(res.Rooms))
	for i, room := range res.Rooms {
		rooms[i] = &model.Room{
			ID:        room.Id,
			UserID:    room.UserId,
			CreatedAt: room.CreatedAt.AsTime().String(),
			UpdatedAt: room.UpdatedAt.AsTime().String(),
		}
	}

	return rooms, nil
}

// GetRoom is the resolver for the getRoom field.
func (r *queryResolver) GetRoom(ctx context.Context, input model.GetRoomInput) (*model.Room, error) {
	res, err := r.RoomClient.GetRoom(ctx, &grpc.GetRoomRequest{
		Id: input.ID,
	})
	if err != nil {
		return nil, err
	}

	return &model.Room{
		ID:        res.Room.Id,
		UserID:    res.Room.UserId,
		CreatedAt: res.Room.CreatedAt.AsTime().String(),
		UpdatedAt: res.Room.UpdatedAt.AsTime().String(),
	}, nil
}
