package check

type VersionChecker interface {
	Version() (string, error)
	ActiveClient() (string, error)
	Health() (string, error)
}
