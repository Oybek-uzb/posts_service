package main

import (
	"context"
	"fmt"
	"net"

	"github.com/Oybek-uzb/posts_service/config"
	postsPostgres "github.com/Oybek-uzb/posts_service/internal/posts/db/postgres"
	"github.com/Oybek-uzb/posts_service/internal/services"
	pbp "github.com/Oybek-uzb/posts_service/pkg/api/posts_service"
	"github.com/Oybek-uzb/posts_service/pkg/client/postgres"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	postgreSQLClient, err := postgres.NewClient(context.TODO(), 3, cfg.Storage)
	if err != nil {
		fmt.Println(fmt.Errorf("%s", err))
	}

	repository := postsPostgres.NewRepository(postgreSQLClient)

	s := grpc.NewServer()
	postsService := services.NewPostsService(repository)
	pbp.RegisterPostsServiceServer(s, postsService)

	listen, err := net.Listen("tcp", cfg.HttpPort)
	if err != nil {
		return
	}
	fmt.Printf("Listening HTTP on %s\n", cfg.HttpPort)

	err = s.Serve(listen)
	if err != nil {
		return
	}

}
