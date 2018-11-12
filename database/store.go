package database

type VersionCheck interface {
	Version() (string, error)
}
