package psql

import (
	"github.com/onkiit/dbcheck/registry"
)

type Store interface {
	GetInfo() (*Postgre, error)
}

func init() {
	registry.RegisterModels()
}
