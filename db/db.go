package db

// The Db interface defines all the public methods provided by the database.
type Db interface {
	Get(key string) (*Record, error)
	Set(key string, value []byte) (*Record, error)
	Delete(key string) (*Record, error)
}
