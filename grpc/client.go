package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func connect() (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	}
	conn, err := grpc.Dial("localhost:8080", opts...)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func ConnectUserServiceClient() (UserServiceClient, error) {
	conn, err := connect()
	if err != nil {
		return nil, err
	}

	return NewUserServiceClient(conn), nil
}

func ConnectTweetServiceClient() (TweetServiceClient, error) {
	conn, err := connect()
	if err != nil {
		return nil, err
	}

	return NewTweetServiceClient(conn), nil
}

func ConnectChatServiceClient() (ChatServiceClient, error) {
	conn, err := connect()
	if err != nil {
		return nil, err
	}

	return NewChatServiceClient(conn), nil
}

func ConnectRoomServiceClient() (RoomServiceClient, error) {
	conn, err := connect()
	if err != nil {
		return nil, err
	}

	return NewRoomServiceClient(conn), nil
}
