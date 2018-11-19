package psql

type Postgre struct {
	Version      string
	ActiveClient int
	Size         []Size
}

type Size struct {
	Name  string
	Usage int64
}
