package storage

type NoopStorage struct {
}

func (n *NoopStorage) Ping() error {
	return nil
}
