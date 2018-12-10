package redis

import (
	"context"
	"fmt"

	"github.com/onkiit/dbinfo"
	rd "github.com/onkiit/dbinfo/redis"

	"github.com/onkiit/dbcheck"
	"github.com/onkiit/dbcheck/registry"

	"github.com/gomodule/redigo/redis"
)

type rediss struct {
	con redis.Conn
}

func (r *rediss) Version() error {
	con := &dbinfo.Conn{
		Con: r.con,
	}

	store := rd.New(con)
	v, err := store.GetVersion(context.Background())
	if err != nil {
		return err
	}

	fmt.Print(v.Version)

	return nil
}

func (r *rediss) ActiveClient() error {
	con := &dbinfo.Conn{
		Con: r.con,
	}

	store := rd.New(con)
	c, err := store.GetActiveClient(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf("active_client(s): %d\n", c.ActiveClient)

	return nil
}

func (r *rediss) Health() error {
	con := &dbinfo.Conn{
		Con: r.con,
	}

	store := rd.New(con)
	h, err := store.GetHealth(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf(" Available Key: %d\n Memory Usage: %s\n Expired Keys: %s\n Evicted Keys: %s\n Slowlog Count: %d\n", h.RedisHealth.AvailableKey, h.RedisHealth.MemoryUsage, h.RedisHealth.ExpiredKeys, h.RedisHealth.EvictedKeys, h.RedisHealth.SlowlogCount)

	fmt.Printf(" Memory Stats\n  Peak Allocated: %d \n  Total Allocated: %d\n  Startup Allocated: %d\n  Peak Percentage: %d\n  Fragmentation: %d\n", h.RedisHealth.MemoryStats.PeakAllocated, h.RedisHealth.MemoryStats.TotalAllowed, h.RedisHealth.MemoryStats.StartupAllocated, h.RedisHealth.MemoryStats.PeakPercentage, h.RedisHealth.MemoryStats.Fragmentation)
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
