package mongo

import (
	"context"
	"fmt"

	"github.com/onkiit/dbinfo"
	mg "github.com/onkiit/dbinfo/mongo"

	"github.com/onkiit/dbcheck"
	"github.com/onkiit/dbcheck/registry"

	"github.com/globalsign/mgo"
)

type mongo struct {
	session *mgo.Session
}

func (m *mongo) Version() error {
	con := &dbinfo.Conn{
		Session: m.session,
	}

	store := mg.New(con)
	v, err := store.GetVersion(context.Background())
	if err != nil {
		return err
	}

	fmt.Println(v.Version)
	return nil
}

func (m *mongo) ActiveClient() error {
	con := &dbinfo.Conn{
		Session: m.session,
	}

	store := mg.New(con)
	c, err := store.GetActiveClient(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf(" active_client(s): %d \n", c.ActiveClient)

	return nil
}

func (m *mongo) Health() error {
	con := &dbinfo.Conn{
		Session: m.session,
	}

	store := mg.New(con)
	h, err := store.GetHealth(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf(" health_status: \n")
	fmt.Printf("  DB: %s\n  Collection: %d\n  Storage Size: %f\n  Indexes: %d\n  Data Size: %f\n", h.MongoHealth.DBName, h.MongoHealth.AvailableCollection, h.MongoHealth.StorageSize, h.MongoHealth.Indexes, h.MongoHealth.DataSize)

	return nil
}

func (m *mongo) Dial(host string) dbcheck.Checker {
	session, err := mgo.Dial(host)
	if err != nil {
		fmt.Println("mongo conn", err)
		return nil
	}

	return &mongo{session: session}
}

func init() {
	registry.Register("mongo", &mongo{})
}
