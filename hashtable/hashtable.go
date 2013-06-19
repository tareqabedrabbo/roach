package hashtable

import (
	"container/list"
	"fmt"
	"hash/fnv"
	"roach/db"
	"sync"
)

const bucketsSize uint32 = 100000

type bucket struct {
	sync.RWMutex
	*list.List
}

type Hashtable struct {
	buckets [bucketsSize]*bucket
}

func (hashtable *Hashtable) Get(key string) (*db.Record, error) {
	var (
		h      = hash(key)
		bucket = hashtable.buckets[h]
	)

	bucket.RLock()
	defer bucket.RUnlock()

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
		bucket = newBucket()
		// does this require locking the hashtable?
		hashtable.buckets[h] = bucket
	}

	bucket.Lock()
	defer bucket.Unlock()

	// find the record and update it
	var r *db.Record
	if e := findInBucket(bucket, key); e == nil {
		r = db.NewRecord(key, value)
		bucket.PushFront(r)
	} else {
		r = e.Value.(*db.Record)
		r.Update(value)
	}
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

	bucket.Lock()
	defer bucket.Unlock()

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

func findInBucket(bucket *bucket, key string) *list.Element {
	for e := bucket.Front(); e != nil; e = e.Next() {
		if r := e.Value.(*db.Record); r.Key() == key {
			return e
		}
	}
	return nil
}

func newBucket() *bucket {
	return &bucket{List: list.New()}
}

func (hashtable *Hashtable) String() string {
	return fmt.Sprintf("{<%d>}", len(hashtable.buckets))
}
