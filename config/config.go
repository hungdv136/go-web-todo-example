package config

import "sync"

type Config interface {
	GetJwtSecretKey() string
	GetDbHost() string
	GetCacheHost() string
}

type todoConfig struct {
	jwtSecretKey string
	databaseHost string
	cacheHost    string
}

func (c *todoConfig) GetJwtSecretKey() string {
	return c.jwtSecretKey
}

func (c *todoConfig) GetDbHost() string {
	return c.databaseHost
}

func (c *todoConfig) GetCacheHost() string {
	return c.cacheHost
}

var instance *todoConfig
var once sync.Once

func GetConfig() Config {
	once.Do(func() {
		instance = &todoConfig{
			jwtSecretKey: "43793aad711f441cb7b69225af6b82d6",
			databaseHost: "root:neogov@123@/todo",
			cacheHost:    "localhost:6379",
		}
	})
	return instance
}
