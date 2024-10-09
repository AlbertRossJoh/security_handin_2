package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	proto "security_handin_2/grpc"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func RegisterOutput(share *proto.Share, toServerId string, creds credentials.TransportCredentials) (proto.ErrorCode, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", toServerId, strconv.Itoa(SERVER_PORT)),
		grpc.WithTransportCredentials(creds))

	if err != nil {
		log.Println("Could not connect to server: ", toServerId)
		return 0, errors.New("Could not connect to server: " + toServerId)
	}

	hospitalServiceClient := proto.NewHospitalServiceClient(conn)

	ack, err := hospitalServiceClient.RegisterOutput(context.Background(), share)

	if err != nil {
		return 0, err
	}

	return ack.ErrorCode, nil
}

func RegisterClient(fromServerId string, toServerId string, creds credentials.TransportCredentials) (proto.ErrorCode, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", toServerId, strconv.Itoa(SERVER_PORT)),
		grpc.WithTransportCredentials(creds))

	if err != nil {
		log.Println("Could not connect to server: ", toServerId)
		return 0, errors.New("Could not connect to server: " + toServerId)
	}

	hospitalServiceClient := proto.NewHospitalServiceClient(conn)

	ack, err := hospitalServiceClient.RegisterClient(context.Background(), &proto.Id{Id: fromServerId})

	if err != nil {
		return 0, err
	}

	return ack.ErrorCode, nil
}
