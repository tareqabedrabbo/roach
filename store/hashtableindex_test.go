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

	log.Printf("==>%+v\n", r)

	if r2 := index.Get(key); r2 != r {
		t.Errorf("Expected %s. Found %v\n", r.Key(), r2)
	}
}
