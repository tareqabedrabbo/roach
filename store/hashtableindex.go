package store

import(
	"hash/fnv"
	"container/list"
	"fmt"
)

const bucketsSize uint32 = 100000

type HashIndex struct {
	buckets [bucketsSize]*list.List
}

func NewHashIndex() *HashIndex {
	return &HashIndex{}
}

func (index *HashIndex) Get(key string) *record {
	var (
		h = hash(key)
		bucket = index.buckets[h]
	)
	for e := bucket.Front(); e != nil; e = e.Next() {
		if r := e.Value.(*record); r.Key() == key {
			return r
		}
	}
	return nil
}

func (index *HashIndex) Set(key string, value []byte) *record {
	h := hash(key)
	bucket := index.buckets[h]

	// find the bucket
	if bucket == nil {
		bucket = list.New()
		index.buckets[h] = bucket
	}

	// find the record and update it
	var r *record
	if e := findInBucket(bucket, key); e == nil {
		r = NewRecord(key, value)		
	} else {
		r = UpdateRecord(e.Value.(*record), value)
		
		// need to guarantee atomicity of update here otherwise the record will disappear before
		// reappearing updated
		bucket.Remove(e)
	}
	
	bucket.PushFront(r)
	return  r
}

func (index *HashIndex) Delete(key string) (*record, error) {
	var (
		h = hash(key)
		bucket = index.buckets[h]
	)
	
	if bucket == nil {
		return nil, fmt.Errorf("key [%s] does not exist", key)
	}

	e := bucket.Front()
	for ; e != nil; e = e.Next() {
		if r := e.Value.(*record); r.Key() == key {
			return bucket.Remove(e).(*record), nil
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
		if r := e.Value.(*record); r.Key() == key {
			return e
		}
	}
	return nil
}

