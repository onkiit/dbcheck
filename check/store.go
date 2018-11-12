package check

type VersionChecker interface {
	Version() (string, error)
}
