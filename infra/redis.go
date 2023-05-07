package infra

import (
	"log"
	"os"

	"github.com/garyburd/redigo/redis"
)

// Conn Redis接続情報
var Conn redis.Conn

func NewRedisConnection() {
	host := os.Getenv("REDIS_HOST")
	port := "6379"

	var err error
	// redis-serverに接続する
	Conn, err = redis.Dial("tcp", host+":"+port)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func NewRedisLocalConnection() {
	host := os.Getenv("REDIS_HOST")
	port := "6379"

	var err error
	// redis-serverに接続する
	Conn, err = redis.Dial("tcp", host+":"+port)
	if err != nil {
		log.Fatal(err)
		return
	}
}
