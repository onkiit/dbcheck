package database

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type Redis struct {
	Host string
}

func (r Redis) Version() (string, error) {
	con, err := redis.Dial("tcp", r.Host)
	if err != nil {
		return "", err
	}

	version, err := redis.String(con.Do("INFO"))
	if err != nil {
		fmt.Println("getting info", err)
		return "", nil
	}

	return version, nil
}

func NewRedis(host string) VersionCheck {
	return &Redis{
		Host: host,
	}
}
