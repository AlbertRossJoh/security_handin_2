package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	proto "security_handin_2/grpc"
	"security_handin_2/shared"
	"security_handin_2/shared/cert"
	"security_handin_2/shared/service"
)

var (
	dockerId      = os.Getenv("HOSTNAME")
	SERVER_PORT   = 6969
	certPath      = "/var/certs/" + dockerId + "-cert.pem"
	keyPath       = "/var/keys/clients/" + dockerId + "-key.pem"
	caCertPath    = "/var/certs/ca/ca-cert.pem"
	serverContext = ServerContext{Id2Int: map[string]int{}}
	outShare      = OutShare{out: 0}
)

const (
	NodesFilename = "nodes"
)

type Server struct {
	proto.UnimplementedShareServiceServer
	id string
}

type ServerContext struct {
	Id2Int map[string]int
}

func main() {
	println("hi from " + dockerId)
	shared.WriteToSharedFile(dockerId, shared.GetPath(NodesFilename))
	config, err := cert.LoadTLSCredentials(certPath, keyPath, caCertPath)
	if err != nil {
		log.Fatal("Could not load cert")
	}
	go StartServer(config)
	time.Sleep(time.Second * 4)

	hospitalId := shared.GetFileContents(dockerId, shared.GetPath("hospital"))[0]
	service.RegisterClient(dockerId, hospitalId, config)
	contents := shared.GetFileContents(dockerId, shared.GetPath(NodesFilename))
	for i, id := range contents {
		serverContext.Id2Int[id] = i
	}

	pctx := PartyContext{
		AmountOfParties: len(contents),
		Prime:           326257,
	}
	sn := rand.Intn(1000)
	log.Printf("secret value: %d", sn)
	secret := NewSecret(sn, pctx)
	outShare.RegisterShare(secret.GetShare(dockerId), dockerId)

	for _, id := range contents {
		if id != dockerId {
			_, err := service.RegisterShare(
				&proto.Share{
					Id:      dockerId,
					Message: int32(secret.GetShare(id)),
				}, id, config)

			if err != nil {
				log.Println("could not dial: " + id)
				log.Println(err.Error())
				continue
			}
			//log.Println(res)
		}
	}
	time.Sleep(time.Second * 4)
	outShare.PrintShare()
	service.RegisterOutput(&proto.Share{Id: dockerId, Message: int32(outShare.out)}, hospitalId, config)
}
