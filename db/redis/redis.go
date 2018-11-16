package redis

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"

	"github.com/onkiit/dbcheck"

	"github.com/gomodule/redigo/redis"
	"github.com/onkiit/dbcheck/registry"
)

type rediss struct {
	con redis.Conn
}

func (r *rediss) redisDo(command string) (string, error) {
	info, err := redis.String(r.con.Do(command))
	if err != nil {
		return "", err
	}
	return info, nil
}

func (r *rediss) getString(info string, prefix string) (string, error) {
	var text string
	reader := bufio.NewReader(bytes.NewBufferString(info))
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}

		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, prefix) {
			text = line
			break
		}
	}
	return text, nil
}

func (r *rediss) Version() error {
	info, err := r.redisDo("INFO")
	if err != nil {
		return err
	}

	version, err := r.getString(info, "redis_version")
	if err != nil {
		return err
	}

	fmt.Println(version)

	return nil
}

func (r *rediss) ActiveClient() error {
	info, err := r.redisDo("INFO")
	if err != nil {
		return err
	}

	client, err := r.getString(info, "connected_clients")
	if err != nil {
		return err
	}

	fmt.Println(client)

	return nil
}

func (r *rediss) Health() error {

	return nil
}

func (r *rediss) Dial(host string) dbcheck.Checker {
	con, err := redis.Dial("tcp", host)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &rediss{con: con}
}

func init() {
	registry.Register("redis", &rediss{})
}
