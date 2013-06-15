package store

import (
	"time"
	"bytes"
	"fmt"
)

type Index interface {
	Get(key string) *record
	// creates a new updated record
	Set(key string, value []byte) *record
	Delete(key string) (*record, error)
}

// immutable structure representing a key/value record
type record struct {
	created int64
	updated int64
	key string
	value bytes.Buffer
}

func NewRecord(key string, value []byte) *record {
	now := time.Now().Unix()
	r := &record{created: now, updated: now, key: key}
	r.value.Write(value)
	return r
}

func UpdateRecord(r *record, newValue []byte) *record {
	now := time.Now().Unix()
	newRecord := &record{created: r.created, updated: now, key: r.key}
	newRecord.value.Write(newValue)
	return newRecord
}

func (r *record) Created() int64 {
	return r.created
}

func (r *record) Updated() int64 {
	return r.updated
}

func (r *record) Key() string {
	return r.key
}

func (r *record) Value() []byte {
	return r.value.Bytes()
}

func (r *record) String() string {
	return fmt.Sprintf("{key: %s created: %d, updated: %d, value: <%v bytes>}", r.key, r.created, r.updated, r.value.Len())
}

