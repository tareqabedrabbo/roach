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
	r := &record{created: now, updated: now, key: key}
	r.value.Write(value)
	return r
}

func (r *record) Update(newValue []byte) {
	now := time.Now().Unix()
	r.value.Reset()
	r.value.Write(newValue)
	r.updated = now
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

