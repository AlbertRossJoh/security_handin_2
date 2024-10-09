package main

import (
	"log"
	"os"
	proto "security_handin_2/grpc"
	"security_handin_2/shared"
	"security_handin_2/shared/cert"
	"sync"
	"time"
)

var (
	dockerId    = os.Getenv("HOSTNAME")
	SERVER_PORT = 6969
	certPath    = "/var/certs/" + dockerId + "-cert.pem"
	keyPath     = "/var/keys/clients/" + dockerId + "-key.pem"
	caCertPath  = "/var/certs/ca/ca-cert.pem"
	aggregate   = Aggregate{sum: 0}
)

type Aggregate struct {
	sum int
	wg  sync.WaitGroup
}

type Server struct {
	proto.UnimplementedHospitalServiceServer
	id string
}

func main() {
	shared.WriteToSharedFile(dockerId, shared.GetPath("hospital"))
	creds, err := cert.LoadTLSCredentials(certPath, keyPath, caCertPath)
	if err != nil {
		log.Fatal("Could not load cert")
	}
	go StartHospitalServer(creds)
	time.Sleep(time.Second * 8)
	aggregate.wg.Wait()
	log.Printf("sum: %d", aggregate.sum)
}
