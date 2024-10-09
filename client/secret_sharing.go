package main

import (
	"log"
	"math/rand"
)

type Secret struct {
	shares []int
}

type OutShare struct {
	out int
}

type PartyContext struct {
	AmountOfParties int
	Prime           int
}

var (
	registeredShareIds = map[string]bool{}
)

func NewSecret(secret int, ctx PartyContext) Secret {
	sum := 0
	var shares []int
	for i := 0; i < ctx.AmountOfParties-1; i++ {
		si := rand.Intn(ctx.Prime - 2)
		sum += si
		shares = append(shares, si)
	}
	shares = append(shares, secret-sum)
	return Secret{
		shares: shares,
	}
}

func (s *Secret) GetShare(id string) int {
	return s.shares[serverContext.Id2Int[id]]
}

func (o *OutShare) RegisterShare(share int, id string) {
	_, ok := registeredShareIds[id]
	if ok {
		return
	}
	o.out += share
	registeredShareIds[id] = true
}

func (o *OutShare) PrintShare() {
	log.Println(o.out)
}
