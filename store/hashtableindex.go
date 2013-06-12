package store

import(
	"hash/fnv"
	"container/list"
	"log"
)

const bucketsSize uint32 = 100000


type HashIndex struct {
	buckets [bucketsSize]list.List
}

func NewHashIndex() *HashIndex {
	return new(HashIndex)
}

func (index *HashIndex) Get(key string) *Record {
	return new(Record)
}

func (index *HashIndex) Set(key string, value []byte) *Record {
	h := hash(key)
	log.Printf("hash = %d\n", hash)
	l := index.buckets[h]
	log.Printf("list = %+v\n", l)
	r := NewRecord(key, value)
	l.PushFront(r)
	log.Printf("list = %+v\n", l)
	return  r
}

func (index *HashIndex) Delete(key string) *Record {
	return new(Record)
}

func hash(key string) uint32 {
	hash := fnv.New32a()
	hash.Write([]byte(key))
	return hash.Sum32() % bucketsSize
}