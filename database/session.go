package database

import (
	"fmt"
	"log"

	"github.com/boj/redistore"
)

func RediStore(size int, network string, host string, port string, password string, secretKey []byte) *redistore.RediStore {
	rdStore, err := redistore.NewRediStore(size, network, fmt.Sprintf("%s:%s", host, port), password, secretKey)
	if err != nil {
		log.Panic(err)
	}

	return rdStore
}
