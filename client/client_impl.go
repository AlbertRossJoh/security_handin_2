package main

import (
	"context"
	"log"
	"net"
	proto "security_handin_2/grpc"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var recievedIds = map[string]bool{}

func StartServer(creds credentials.TransportCredentials) {
	server := &Server{
		id: dockerId,
	}

	listener, err := net.Listen("tcp", ":"+strconv.Itoa(SERVER_PORT))
	if err != nil {
		log.Fatalf("Could not create the server %v", err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))
	proto.RegisterShareServiceServer(grpcServer, server)
	serveError := grpcServer.Serve(listener)

	if serveError != nil {
		log.Fatal("Could not serve listener")
	}
}

func (s *Server) Test(ctx context.Context, in *proto.EmptyArg) (*proto.Ack, error) {
	return &proto.Ack{
		ErrorCode: proto.ErrorCode_SUCCESS,
	}, nil
}

func (s *Server) RegisterShare(ctx context.Context, in *proto.Share) (*proto.Ack, error) {
	shareChan <- in

	return &proto.Ack{
		ErrorCode: proto.ErrorCode_SUCCESS,
	}, nil
}
