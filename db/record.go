package db

import (
	"bytes"
	"fmt"
	"time"
)

// Record is a data structure representing a key/value entry in the database.
type Record struct {
	created int64
	updated int64
	key     string
	value   bytes.Buffer
}

func NewRecord(key string, value []byte) *Record {
	now := time.Now().Unix()
	r := &Record{created: now, updated: now, key: key}
	r.value.Write(value)
	return r
}

func UpdateRecord(r *Record, newValue []byte) *Record {
	now := time.Now().Unix()
	newRecord := &Record{created: r.created, updated: now, key: r.key}
	newRecord.value.Write(newValue)
	return newRecord
}

func (r *Record) Created() int64 {
	return r.created
}

func (r *Record) Updated() int64 {
	return r.updated
}

func (r *Record) Key() string {
	return r.key
}

func (r *Record) Value() []byte {
	return r.value.Bytes()
}

func (r *Record) String() string {
	return fmt.Sprintf("{key: %s created: %d, updated: %d, value: <%v bytes>}", r.key, r.created, r.updated, r.value.Len())
}
