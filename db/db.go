package db

type Db interface {
	Get(key string) *Record
	// creates a new updated record
	Set(key string, value []byte) *Record
	Delete(key string) (*Record, error)
}
