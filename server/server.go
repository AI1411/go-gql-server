package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/rs/cors"

	generated "github.com/AI1411/go-gql-server/graph"
	grpcClient "github.com/AI1411/go-gql-server/graph/grpc"
	"github.com/AI1411/go-gql-server/graph/resolver"
)

const (
	defaultGraphqlPort = "8081"
	defaultClientPort  = "3000"
)

func Server() {
	port := os.Getenv("STAR_GRAPHQL_PORT")

	if port == "" {
		port = defaultGraphqlPort
	}

	userClient, err := grpcClient.ConnectUserServiceClient()
	if err != nil {
		log.Fatalf("failed to connect to user server: %v", err)
	}

	tweetClient, err := grpcClient.ConnectTweetServiceClient()
	if err != nil {
		log.Fatalf("failed to connect to tweet server: %v", err)
	}

	chatClient, err := grpcClient.ConnectChatServiceClient()
	if err != nil {
		log.Fatalf("failed to connect to chat server: %v", err)
	}

	roomClient, err := grpcClient.ConnectRoomServiceClient()
	if err != nil {
		log.Fatalf("failed to connect to room server: %v", err)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{
		UserClient:  userClient,
		TweetClient: tweetClient,
		ChatClient:  chatClient,
		RoomClient:  roomClient,
	}}))

	if os.Getenv("STAR_CLIENT_PORT") == "" {
		os.Setenv("STAR_CLIENT_PORT", defaultClientPort)
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{fmt.Sprintf("http://localhost:%s", os.Getenv("STAR_CLIENT_PORT"))},
		AllowCredentials: true,
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	http.Handle("/query", c.Handler(srv))

	if err = http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
		return
	}
}
