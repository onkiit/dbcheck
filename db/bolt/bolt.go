package bolt

import (
	"context"
	"fmt"
	"time"

	"github.com/onkiit/dbinfo"
	bl "github.com/onkiit/dbinfo/bolt"

	"github.com/boltdb/bolt"
	"github.com/onkiit/dbcheck"
	"github.com/onkiit/dbcheck/registry"
)

type boltdb struct {
	db *bolt.DB
}

func (b *boltdb) Version() error {
	fmt.Println("Bolt does not support version checking.")
	return nil
}

func (b *boltdb) ActiveClient() error {
	con := &dbinfo.Conn{
		BoltDB: b.db,
	}

	store := bl.New(con)
	c, err := store.GetActiveClient(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf("Transaction: %d\n", c.ActiveClient)
	return nil
}

func (b *boltdb) Health() error {
	con := &dbinfo.Conn{
		BoltDB: b.db,
	}

	store := bl.New(con)
	h, err := store.GetHealth(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf("health information: \n Available Bucket: %d\n Available Key: %d\n", h.BoltHealth.NumberOfBucket, h.BoltHealth.NumberOfKey)
	return nil
}

func (b *boltdb) Dial(host string) dbcheck.Checker {
	db, err := bolt.Open(host, 0600, &bolt.Options{Timeout: time.Second * 5})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &boltdb{db: db}
}

func init() {
	registry.Register("bolt", &boltdb{})
}
