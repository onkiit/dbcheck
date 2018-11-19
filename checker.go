package dbcheck

type Checker interface {
	Version() error
	ActiveClient() error
	Health() error
}

type Dialer interface {
	Dial(string) Checker
}
