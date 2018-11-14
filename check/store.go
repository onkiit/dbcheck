package check

type VersionChecker interface {
	Version() error
}

type ClientChecker interface {
	ActiveClient() error
}

type HealthChecker interface {
	Health() error
}

type Dialer interface {
	Dial() error
}

type DBChecker interface {
	GetInfo() error
}
