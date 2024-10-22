package main

import (
	"log"
	"math/rand"
	"os"
	"security_handin_2/shared"
	"security_handin_2/shared/cert"
	"security_handin_2/shared/service"

	proto "security_handin_2/grpc"

	"github.com/google/uuid"
	"google.golang.org/grpc/credentials"
)

var (
	dockerId          = os.Getenv("HOSTNAME")
	SERVER_PORT       = 6969
	serverCertPath    = "/var/certs/" + dockerId + "-server-cert.pem"
	clientCertPath    = "/var/certs/" + dockerId + "-client-cert.pem"
	serverKeyPath     = "/var/keys/clients/" + dockerId + "-server-key.pem"
	clientKeyPath     = "/var/keys/clients/" + dockerId + "-client-key.pem"
	caCertPath        = "/var/certs/ca/ca-cert.pem"
	serverContext     = ServerContext{Id2Int: map[string]int{}}
	outShare          = OutShare{out: 0}
	shareChan         = make(chan *proto.Share, 1000)
	registeredIds     = make(map[string]string)
	clientNameToCreds = make(map[string]credentials.TransportCredentials)
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
	log.Println("hi from: " + dockerId)
	// shared.WriteToSharedFile(dockerId, shared.GetPath(NodesFilename))
	serverCreds, err := cert.LoadTLSServerCredentials(serverCertPath, serverKeyPath, caCertPath)
	if err != nil {
		log.Panicln(err.Error())
		log.Fatal("Could not load server cert")
	}

	go StartServer(serverCreds)

	hospitalId := shared.GetFileContents(dockerId, shared.GetPath("hospital"))[0]
	hospitalCreds := cert.LoadTLSClientCredentials(clientCertPath, clientKeyPath, caCertPath, hospitalId)

	service.WaitForHospitalServiceStart([]string{hospitalId}, hospitalCreds)
	_, err = service.RegisterClient(dockerId, hospitalId, hospitalCreds)
	if err != nil {
		log.Panicln(err.Error())
	}
	contents := shared.GetFileContents(dockerId, shared.GetPath(NodesFilename))
	for _, id := range contents {
		registeredIds[id] = ""
		clientNameToCreds[id] = cert.LoadTLSClientCredentials(clientCertPath, clientKeyPath, caCertPath, id)
	}
	service.WaitForClientServiceStart(contents, clientNameToCreds)

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
					Guid:    uuid.NewString(),
				}, id, clientNameToCreds[id])
			if err != nil {
				log.Println("could not dial: " + id)
				log.Panic(err.Error())
				continue
			}
		}
	}

	for i := 0; i < len(contents)-1; i++ {
		in := <-shareChan
		// for idempotence
		guid, ok := registeredIds[in.Id]
		if ok && in.Guid != guid {
			outShare.RegisterShare(int(in.Message), in.Id)
			registeredIds[in.Id] = guid
		} else {
			if !ok {
				log.Panic("unknown id")
			}
		}
	}

	// patiently wait for the other clients
	// time.Sleep(time.Second * 4)
	outShare.PrintShare()
	service.RegisterOutput(&proto.Share{Id: dockerId, Message: int32(outShare.out), Guid: uuid.NewString()}, hospitalId, hospitalCreds)
}
