package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	proto "security_handin_2/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const SERVER_PORT = 6969

func Test(serverId string, creds credentials.TransportCredentials) (proto.ErrorCode, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", serverId, strconv.Itoa(SERVER_PORT)),
		grpc.WithTransportCredentials(creds))

	if err != nil {
		log.Println("Could not connect to server: ", serverId)
		return 0, errors.New("Could not connect to server: " + serverId)
	}

	shareServiceClient := proto.NewShareServiceClient(conn)

	ack, err := shareServiceClient.Test(context.Background(), &proto.EmptyArg{})

	if err != nil {
		return 0, err
	}

	return ack.ErrorCode, nil
}

func RegisterShare(share *proto.Share, toServerId string, creds credentials.TransportCredentials) (proto.ErrorCode, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", toServerId, strconv.Itoa(SERVER_PORT)),
		grpc.WithTransportCredentials(creds))

	if err != nil {
		log.Println("Could not connect to server: ", toServerId)
		return 0, errors.New("Could not connect to server: " + toServerId)
	}

	shareServiceClient := proto.NewShareServiceClient(conn)

	ack, err := shareServiceClient.RegisterShare(context.Background(), share)

	if err != nil {
		return 0, err
	}

	return ack.ErrorCode, nil
}
