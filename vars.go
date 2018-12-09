package main

import (
	"flag"

	"github.com/go-redis/redis"

	"github.com/dgraph-io/badger"
)

var (
	flagListenAddr = flag.String("listen", ":8044", "the listen address")
	flagRedisAddr  = flag.String("redis", "redis://localhost/2", "redis address")
	// flagStorage    = flag.String("storage", filepath.Join(filepath.Dir(os.Args[0]), "utree"), "the storage directory")
)

var (
	db        *badger.DB
	redisConn *redis.Client
)

var (
	redisTreePrefix = "utree:tree:"
)

var (
	banner = `

     _______            
    |__   __|           
  _   _| |_ __ ___  ___ 
 | | | | | '__/ _ \/ _ \
 | |_| | | | |  __/  __/
  \__,_|_|_|  \___|\___|
                                         
	
welcome to the uflarians tree handler, it helps you
write a high perfomant tree based features like teams, categories, ... etc
you can use redis directly, but this to just unifiy or standardize the concept.
	
	`
)
