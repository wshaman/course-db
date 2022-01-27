package datasource

type mockDB struct {
}

func NewMock() DS {
	return &mockDB{}
}
