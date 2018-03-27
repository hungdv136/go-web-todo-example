package configs

import "sync"

type Config interface {
	GetJwtSecretKey() string
	GetDbHost() string
}

type todoConfig struct {
	jwtSecretKey string
	databaseHost string
}

func (c *todoConfig) GetJwtSecretKey() string {
	return c.jwtSecretKey
}

func (c *todoConfig) GetDbHost() string {
	return c.databaseHost
}

var instance *todoConfig
var once sync.Once

func GetConfig() Config {
	once.Do(func() {
		instance = &todoConfig{
			jwtSecretKey: "43793aad711f441cb7b69225af6b82d6",
			databaseHost: "root:neogov@123@/todo",
		}
	})
	return instance
}
