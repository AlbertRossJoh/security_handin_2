package service

import (
	"context"

	proto "security_handin_2/grpc"

	"google.golang.org/grpc/credentials"
)

func Test(serverId string, creds credentials.TransportCredentials) (proto.ErrorCode, error) {
	conn := getClientConn(serverId, creds)

	shareServiceClient := proto.NewShareServiceClient(conn)

	ack, err := shareServiceClient.Test(context.Background(), &proto.EmptyArg{})
	if err != nil {
		return 0, err
	}

	return ack.ErrorCode, nil
}

func RegisterShare(share *proto.Share, toServerId string, creds credentials.TransportCredentials) (proto.ErrorCode, error) {
	conn := getClientConn(toServerId, creds)

	shareServiceClient := proto.NewShareServiceClient(conn)

	ack, err := shareServiceClient.RegisterShare(context.Background(), share)
	if err != nil {
		return 0, err
	}

	return ack.ErrorCode, nil
}

func WaitForClientServiceStart(servers []string, creds map[string]credentials.TransportCredentials) {
	for {
		count := len(servers)
		for _, serverId := range servers {
			if _, err := Test(serverId, creds[serverId]); err == nil {
				count--
			}
		}
		if count == 0 {
			break
		}
	}
}
