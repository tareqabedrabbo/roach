package db

type Db interface {
	Get(key string) (*Record, error)
	Set(key string, value []byte) (*Record, error)
	Delete(key string) (*Record, error)
}
