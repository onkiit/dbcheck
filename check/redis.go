package check

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type rediss struct {
	host string
}

func (r rediss) Version() (string, error) {
	con, err := redis.Dial("tcp", r.host)
	if err != nil {
		return "", err
	}

	defer con.Close()

	version, err := redis.String(con.Do("INFO"))
	if err != nil {
		fmt.Println("getting info", err)
		return "", nil
	}

	return version, nil
}

func (r rediss) ActiveClient() (string, error) {
	return "", nil
}

func (r rediss) Health() (string, error) {
	return "", nil
}

func NewRedis(host string) VersionChecker {
	return rediss{
		host: host,
	}
}
