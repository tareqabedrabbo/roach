package hashtable

import (
	"fmt"
	"log"
	"roach/db"
	"testing"
	"time"
)

func TestHash(t *testing.T) {
	var (
		key = "mykey"
		h   = hash(key)
	)

	log.Printf("hash for key [%s] = %d\n", key, h)
	for i := 0; i < 5; i++ {
		if h2 := hash(key); h2 != h {
			t.Errorf("Hash error. Expected %d, found %d\n", h, h2)
		}
	}
}

func TestSet(t *testing.T) {
	var db, key, value = initData()

	if r, _ := db.Set(key, value); r.Created() == 0 || r.Key() != key {
		t.Error("failed!")
	}
}

func TestSetUpdate(t *testing.T) {
	var db, key, value = initData()
	valLen := 20
	r1, _ := db.Set(key, value)

	fmt.Printf("before sleep [%+v]\n", r1)
	time.Sleep(1 * time.Second)

	r2, _ := db.Set(key, make([]byte, valLen))
	fmt.Printf("after sleep [%+v]\n", r2)

	if c1, c2 := r1.Created(), r2.Created(); c1 != c2 {
		t.Errorf("Expected creation time [%d]. Found [%d].\n", c1, c2)
	}

	if u1, u2 := r1.Updated(), r2.Updated(); u2 <= u1 {
		t.Errorf("Expected updated time more recent than [%d]. Found [%d].\n", u1, u2)
	}

	if newLen := len(r2.Value()); newLen != valLen {
		t.Errorf("Expected new value length to be [%d]. Found [%d].", valLen, newLen)
	}

}

func TestGet(t *testing.T) {
	var hashtable, key, value = initData()

	r, _ := hashtable.Set(key, value)
	r2, _ := hashtable.Get(key)
	if r2 != r {
		t.Errorf("Expected %s. Found %v\n", r.Key(), r2)
	}
}

func TestDeleteExisiting(t *testing.T) {
	var hashtable, key, value = initData()

	hashtable.Set(key, value)

	if r, _ := hashtable.Delete(key); r == nil || r.Key() != key {
		t.Errorf("Error in deleting key [%s]\n", key)
	}
}

func initData() (db db.Db, key string, value []byte) {
	return &Hashtable{}, "mykey", make([]byte, 10)
}
