package cache

import (
	"fmt"
	"todo/config"

	"github.com/mediocregopher/radix.v2/redis"

	"github.com/mediocregopher/radix.v2/pool"
)

var db *pool.Pool

func Init() *pool.Pool {
	var err error
	db, err = pool.New("tcp", config.GetConfig().GetCacheHost(), 10)
	if err != nil {
		fmt.Println("cache init error: ", err)
	}

	return db
}

func Get() *redis.Client {
	cli, err := db.Get()
	if err != nil {
		panic(err)
	}
	return cli
}

func Put(client *redis.Client) {
	db.Put(client)
}
