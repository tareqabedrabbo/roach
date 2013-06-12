package store

import(
	"testing"
)

func TestSet(t *testing.T) {
    var index Index = &HashIndex{}
    key := "mykey"
    value := make([]byte, 10)
    if r := index.Set(key, value); r.created == 0 || r.key != key {
    	t.Error("failed!")
    }
}
