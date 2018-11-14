package check

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"

	"github.com/gomodule/redigo/redis"
)

type rediss struct {
	con  redis.Conn
	host string
}

func (r *rediss) info() (string, error) {
	info, err := redis.String(r.con.Do("INFO"))
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
	info, err := r.info()
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
	info, err := r.info()
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

func (r *rediss) Dial() error {
	con, err := redis.Dial("tcp", r.host)
	if err != nil {
		fmt.Println(err)
		return err
	}
	r.con = con

	return nil
}

func (r *rediss) GetInfo() error {
	if err := r.Dial(); err != nil {
		fmt.Println(err)
		return err
	}

	if err := r.Version(); err != nil {
		fmt.Println(err)
		return err
	}

	if err := r.ActiveClient(); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func NewRedis(host string) DBChecker {
	return &rediss{
		host: host,
	}
}
