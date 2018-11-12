package check

import (
	"fmt"

	"github.com/globalsign/mgo"
)

type Mongo struct {
	Host string
}

func (m Mongo) Version() (string, error) {
	session, err := mgo.Dial(m.Host)
	if err != nil {
		fmt.Println("mongo conn", err)
		return "", err
	}

	buildInfo, err := session.BuildInfo()
	if err != nil {
		fmt.Println("getting build info", err)
		return "", err
	}

	version := fmt.Sprintf(" MongoDB\n db version %s \n git version %s \n OpenSSL version %s \n", buildInfo.Version, buildInfo.GitVersion, buildInfo.OpenSSLVersion)

	return version, nil
}

func NewMongo(host string) VersionCheck {
	return &Mongo{
		Host: host,
	}
}
