package main

import (
	"fmt"
	"github.com/maxim-shishkov/books/pkg/api"
	"github.com/maxim-shishkov/books/pkg/config"
	"github.com/maxim-shishkov/books/pkg/database"
	"google.golang.org/grpc"
	"log"
	"net"
)

func init() {
	config.Init()
}

func main() {
	err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	srv := &api.GRPCServer{}
	api.RegisterGetterServer(s, srv)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Common.ServerPort))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server run", config.Common.ServerPort)
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
