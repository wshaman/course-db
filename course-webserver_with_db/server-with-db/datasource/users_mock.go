package datasource

func (m *mockDB) GetUsers() ([]string, error) {
	return []string{"James Alan Hetfield", "Till Lindemann", "Pär Sundström"}, nil
}
