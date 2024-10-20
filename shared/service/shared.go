package service

import (
	"fmt"
	"log"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const SERVER_PORT = 6969

func getClientConn(toServerId string, creds credentials.TransportCredentials) *grpc.ClientConn {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", toServerId, strconv.Itoa(SERVER_PORT)),
		grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err.Error())
	}
	return conn
}
