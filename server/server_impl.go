package main

import (
	"context"
	"log"
	"net"
	"strconv"

	proto "security_handin_2/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var recievedIds = map[string]bool{}

func StartHospitalServer(creds credentials.TransportCredentials) {
	server := &Server{
		id: dockerId,
	}

	listener, err := net.Listen("tcp", ":"+strconv.Itoa(SERVER_PORT))
	if err != nil {
		log.Fatalf("Could not create the server %v", err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))
	proto.RegisterHospitalServiceServer(grpcServer, server)
	serveError := grpcServer.Serve(listener)

	if serveError != nil {
		log.Fatal("Could not serve listener")
	}
}

func (s *Server) RegisterOutput(ctx context.Context, in *proto.Share) (*proto.Ack, error) {
	used, ok := recievedIds[in.Id]
	if nonce.Register(in.Guid) && ok && !used {
		aggregate.sum += int(in.Message)
		recievedIds[in.Id] = true
		aggregate.wg.Done()
	}
	return &proto.Ack{
		ErrorCode: proto.ErrorCode_SUCCESS,
	}, nil
}

func (s *Server) RegisterClient(ctx context.Context, in *proto.Id) (*proto.Ack, error) {
	aggregate.wg.Add(1)
	clientReg <- true
	recievedIds[in.Id] = false
	return &proto.Ack{
		ErrorCode: proto.ErrorCode_SUCCESS,
	}, nil
}

func (s *Server) Test(ctx context.Context, in *proto.EmptyArg) (*proto.Ack, error) {
	return &proto.Ack{
		ErrorCode: proto.ErrorCode_SUCCESS,
	}, nil
}
