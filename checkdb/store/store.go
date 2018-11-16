package store

type Checker interface {
	Version() error
	ActiveClient() error
	Health() error
}

type Dialer interface {
	Dial(string) Checker
}

var Dials = make(map[string]DialFactory)

type DialFactory func() Dialer

func Register(name string, cf DialFactory) {
	Dials[name] = cf
}
