package check

type VersionChecker interface {
	Version() error
}

type ClientChecker interface {
	ActiveClient() error
}

type Dialer interface {
	Dial() error
}
