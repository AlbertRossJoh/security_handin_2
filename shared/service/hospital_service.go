package service

import (
	"context"
	proto "security_handin_2/grpc"
	"time"

	"google.golang.org/grpc/credentials"
)

func RegisterOutput(share *proto.Share, toServerId string, creds credentials.TransportCredentials) (proto.ErrorCode, error) {
	conn := getClientConn(toServerId, creds)

	hospitalServiceClient := proto.NewHospitalServiceClient(conn)

	ack, err := hospitalServiceClient.RegisterOutput(context.Background(), share)
	if err != nil {
		return 0, err
	}

	return ack.ErrorCode, nil
}

func RegisterClient(fromServerId string, toServerId string, creds credentials.TransportCredentials) (proto.ErrorCode, error) {
	conn := getClientConn(toServerId, creds)

	hospitalServiceClient := proto.NewHospitalServiceClient(conn)

	ack, err := hospitalServiceClient.RegisterClient(context.Background(), &proto.Id{Id: fromServerId})
	if err != nil {
		return 0, err
	}

	return ack.ErrorCode, nil
}

func HospitalTest(serverId string, creds credentials.TransportCredentials) (proto.ErrorCode, error) {
	conn := getClientConn(serverId, creds)

	shareServiceClient := proto.NewHospitalServiceClient(conn)

	ack, err := shareServiceClient.Test(context.Background(), &proto.EmptyArg{})
	if err != nil {
		return 0, err
	}

	return ack.ErrorCode, nil
}

func WaitForHospitalServiceStart(servers []string, creds credentials.TransportCredentials) {
	for {
		count := len(servers)
		for _, serverId := range servers {
			if _, err := HospitalTest(serverId, creds); err == nil {
				count--
			}
			time.Sleep(time.Second)
		}
		if count == 0 {
			break
		}
	}
}
