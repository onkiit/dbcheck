package cassandra

import (
	"context"
	"fmt"

	"github.com/onkiit/dbinfo"
	cs "github.com/onkiit/dbinfo/cassandra"

	"github.com/gocql/gocql"
	"github.com/onkiit/dbcheck"
	"github.com/onkiit/dbcheck/registry"
)

type cassandra struct {
	session *gocql.Session
}

func (c *cassandra) Version() error {
	con := &dbinfo.Conn{
		CQLSession: c.session,
	}
	store := cs.New(con)
	ver, err := store.GetVersion(context.Background())
	if err != nil {
		return err
	}
	fmt.Println(ver.Version)
	return nil
}

func (c *cassandra) ActiveClient() error {
	fmt.Println("Cassandra does not provided active_client checking.")
	return nil
}

func (c *cassandra) Health() error {
	con := &dbinfo.Conn{
		CQLSession: c.session,
	}
	store := cs.New(con)
	h, err := store.GetHealth(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("health information: \n")
	fmt.Printf(" ID\t\t : %s\n Gossip Active   : %s\n Thrift Active   : %s\n Native Transport: %s\n Load\t\t : %s\n Generation No   : %s\n Uptime\t\t : %s\n", h.CassandraHealth.ID, h.CassandraHealth.GossipActive, h.CassandraHealth.ThriftActive, h.CassandraHealth.NativeTransport, h.CassandraHealth.Load, h.CassandraHealth.GenerationNo, h.CassandraHealth.Uptime)
	return nil
}

func (c *cassandra) Dial(host string) dbcheck.Checker {
	cluster := gocql.NewCluster(host)
	if cluster == nil {
		fmt.Println("no cluster")
		return nil
	}

	session, err := cluster.CreateSession()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &cassandra{session: session}
}

func init() {
	registry.Register("cassandra", &cassandra{})
}
