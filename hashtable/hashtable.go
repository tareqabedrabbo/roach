package hashtable

import (
	"container/list"
	"fmt"
	"hash/fnv"
	"roach/db"
)

const bucketsSize uint32 = 100000

type Hashtable struct {
	buckets [bucketsSize]*list.List
}

func NewHashtable() *Hashtable {
	return &Hashtable{}
}

func (hashtable *Hashtable) Get(key string) (*db.Record, error) {
	var (
		h      = hash(key)
		bucket = hashtable.buckets[h]
	)
	for e := bucket.Front(); e != nil; e = e.Next() {
		if r := e.Value.(*db.Record); r.Key() == key {
			return r, nil
		}
	}
	return nil, nil
}

func (hashtable *Hashtable) Set(key string, value []byte) (*db.Record, error) {
	h := hash(key)
	bucket := hashtable.buckets[h]

	// find the bucket
	if bucket == nil {
		bucket = list.New()
		hashtable.buckets[h] = bucket
	}

	// find the record and update it
	var r *db.Record
	if e := findInBucket(bucket, key); e == nil {
		r = db.NewRecord(key, value)
	} else {
		r = db.UpdateRecord(e.Value.(*db.Record), value)

		// need to guarantee atomicity of update here otherwise the record will disappear before
		// reappearing updated
		bucket.Remove(e)
	}

	bucket.PushFront(r)
	return r, nil
}

func (hashtable *Hashtable) Delete(key string) (*db.Record, error) {
	var (
		h      = hash(key)
		bucket = hashtable.buckets[h]
	)

	if bucket == nil {
		return nil, fmt.Errorf("key [%s] does not exist", key)
	}

	e := bucket.Front()
	for ; e != nil; e = e.Next() {
		if r := e.Value.(*db.Record); r.Key() == key {
			return bucket.Remove(e).(*db.Record), nil
		}
	}

	return nil, fmt.Errorf("key [%s] does not exist", key)
}

func hash(key string) uint32 {
	hash := fnv.New32a()
	hash.Write([]byte(key))
	return hash.Sum32() % bucketsSize
}

func findInBucket(bucket *list.List, key string) *list.Element {
	for e := bucket.Front(); e != nil; e = e.Next() {
		if r := e.Value.(*db.Record); r.Key() == key {
			return e
		}
	}
	return nil
}
