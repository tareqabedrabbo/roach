package store

import(
	"hash/fnv"
	"container/list"
	"log"
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
	log.Printf("hash for key %s = %d . Bucket's length = %d\n", key, h, bucket.Len())
	for e := bucket.Front(); e != nil; e = e.Next() {
		if r := e.Value.(*record); r.Key() == key {
			return r
		}
	}
	return nil
}

func (index *HashIndex) Set(key string, value []byte) *record {
	h := hash(key)
	log.Printf("hash for key [%s] = %d\n", key, h)
	l := index.buckets[h]
	if l == nil {
		l = list.New()
		index.buckets[h] = l
	}
	log.Printf("list = %+v\n", l)
	r := NewRecord(key, value)
	l.PushFront(r)
	log.Printf("list = %+v\n", l)
	return  r
}

func (index *HashIndex) Delete(key string) *record {
	return new(record)
}

func hash(key string) uint32 {
	hash := fnv.New32a()
	hash.Write([]byte(key))
	return hash.Sum32() % bucketsSize
}