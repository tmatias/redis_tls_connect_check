package main

import (
	"fmt"
	"log"
	"strconv"

	"os"

	"github.com/gomodule/redigo/redis"
)

func mustBool(s string) bool {
	b, err := strconv.ParseBool(s)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func main() {
	addr := os.Getenv("REDIS_ADDR")
	pass := os.Getenv("REDIS_PASSWORD")
	port := os.Getenv("REDIS_PORT")
	tls := mustBool(os.Getenv("REDIS_USE_TLS"))

	addr = fmt.Sprintf("%s:%s", addr, port)

	log.Println("checking connection...")

	_, err := redis.Dial("tcp", addr, redis.DialPassword(pass), redis.DialUseTLS(tls))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("successfully connected to %s\n", addr)

	select {}
}
