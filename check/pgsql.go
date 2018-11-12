package check

import "database/sql"

type psql struct {
	host string
}

func (p psql) Version() (string, error) {
	db, err := sql.Open("postgres", p.host)
	if err != nil {
		return "", err
	}

	defer db.Close()

	var version string
	_ = db.QueryRow("SELECT version()").Scan(&version)

	return version, nil
}

func (p psql) ActiveClient() (string, error) {
	return "", nil
}

func (p psql) Health() (string, error) {
	return "", nil
}

func NewPsql(host string) VersionChecker {
	return psql{
		host: host,
	}
}
