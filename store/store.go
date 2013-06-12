package store

import (
	"time"
	"bytes"
)

type Index interface {
	Get(key string) *record
	Set(key string, value []byte) *record
	Delete(key string) *record
}

type record struct {
	created int64
	updated int64
	key string
	value bytes.Buffer
}

func NewRecord(key string, value []byte) *record {
	now := time.Now().Unix()
	r := &Record{created: now, updated: now, key: key}
	r.value.Write(value)
	return r
}

func (r *record) Update(newValue []byte) {
	now := time.Now().Unix()
	r.value.Reset()
	r.value.Write(newValue)
	r.updated = now
}

func (r *record) created() int64 {
	return r.created
}

func (r *record) updated() int64 {
	return r.updated
}

func (r *record) key() string {
	return r.key
}

func (r *record) value() []byte {
	return r.value.Bytes()
}

