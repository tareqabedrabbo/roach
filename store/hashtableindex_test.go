package store

import(
	"testing"
	"log"
)

func TestHash(t *testing.T) {
	var (
		key = "mykey"
		h = hash(key)
	)

	log.Printf("hash for key [%s] = %d\n", key, h)
	for i := 0; i < 5; i++ {
		if h2 := hash(key); h2 != h {
			t.Errorf("Hash error. Expected %d, found %d\n", h, h2)
		}
	}
}

func TestSet(t *testing.T) {
    index  := &HashIndex{}
    key := "mykey"
    value := make([]byte, 10)
    if r := index.Set(key, value); r.created == 0 || r.key != key {
    	t.Error("failed!")
    }
}

func TestGet(t *testing.T) {
	var (
		index = &HashIndex{}
		key = "mykey"
		value = make([]byte, 10)
		r = index.Set(key, value)
	)

	if r2 := index.Get(key); r2 != r {
		t.Errorf("Expected %s. Found %v\n", r.Key(), r2)
	}
}

func TestDeleteExisiting(t *testing.T) {
	var index, key, value = initData()
	index.Set(key, value)
	if r, _ := index.Delete(key); r == nil || r.Key() != key {
		t.Errorf("Error in deleting key [%s]\n", key)
	}
}

func initData() (index Index, key string, value []byte) {
	return &HashIndex{}, "mykey", make([]byte, 10)
}
