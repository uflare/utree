package main

import (
	"flag"
	"log"

	"github.com/fatih/color"
	"github.com/go-redis/redis"
)

func init() {
	flag.Parse()

	ro, err := redis.ParseURL(*flagRedisAddr)
	if err != nil {
		log.Fatal(color.RedString("[redis] - %s", err.Error()))
	}

	redisConn = redis.NewClient(ro)
	if _, err := redisConn.Ping().Result(); err != nil {
		log.Fatal(color.RedString("[redis] - %s", err.Error()))
	}
}
