package datasource

type DS interface {
	GetUsers() ([]string, error)
}
