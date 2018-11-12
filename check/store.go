package check

type VersionCheck interface {
	Version() (string, error)
}
