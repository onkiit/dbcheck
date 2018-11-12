package check

import "database/sql"

type psql struct {
	Host string
}

func (p psql) Version() (string, error) {
	db, err := sql.Open("postgres", p.Host)
	if err != nil {
		return "", err
	}

	defer db.Close()

	var version string
	_ = db.QueryRow("SELECT version()").Scan(&version)

	return version, nil
}

func NewPsql(host string) VersionChecker {
	return psql{
		Host: host,
	}
}
