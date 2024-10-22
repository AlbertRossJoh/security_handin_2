package main

import (
	"log"
	"os"
	proto "security_handin_2/grpc"
	"security_handin_2/shared/cert"
	"sync"
)

var (
	dockerId    = os.Getenv("HOSTNAME")
	SERVER_PORT = 6969
	certPath    = "/var/certs/" + dockerId + "-cert.pem"
	keyPath     = "/var/keys/clients/" + dockerId + "-key.pem"
	caCertPath  = "/var/certs/ca/ca-cert.pem"
	aggregate   = Aggregate{sum: 0}
	clientReg   = make(chan bool, 1000)
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
	log.Println("hi from:", dockerId)
	creds, err := cert.LoadTLSServerCredentials(certPath, keyPath, caCertPath)
	if err != nil {
		log.Fatal("Could not load cert")
	}

	go StartHospitalServer(creds)

	// Waits for the first register so we don't go directly to the empty waitgroup
	<-clientReg

	aggregate.wg.Wait()
	log.Printf("sum: %d", aggregate.sum)
}
