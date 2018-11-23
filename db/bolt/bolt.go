package bolt

import (
	"fmt"
	"time"

	"github.com/boltdb/bolt"
	"github.com/onkiit/dbcheck"
	"github.com/onkiit/dbcheck/registry"
)

type boltdb struct {
	db *bolt.DB
}

func (b *boltdb) Version() error {
	fmt.Println("Bolt does not support version cheking.")
	return nil
}

func (b *boltdb) ActiveClient() error {

	return nil
}

func (b *boltdb) Health() error {
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
