package store

import(
	"testing"
	"log"
	"bytes"
)

func TestCreateRecord(t *testing.T) {
	var (	
			key = "mykey"
			value = make([]byte, 10)
			r = NewRecord(key, value)

		)

	log.Printf("created record %+v\n", r)

	if created, updated := r.Created(), r.Updated(); created == 0 || updated == 0 || created != updated {
		t.Errorf("Time was not initialised correctly: %+v\n", r)
	}

	if actualKey := r.Key(); actualKey != key {
		t.Errorf("Expected key [%s]. Found key [%s]", key, actualKey)
	}

	if actualValue := r.Value(); !bytes.Equal(actualValue, value) {
		t.Errorf("Expected value [%v]. Found value [%v]", value, actualValue)
	}
}