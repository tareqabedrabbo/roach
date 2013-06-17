package db

import (
	"bytes"
	"log"
	"testing"
)

func TestCreateRecord(t *testing.T) {
	var (
		key   = "mykey"
		value = make([]byte, 10)
		r     = NewRecord(key, value)
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

func TestUpdateRecord(t *testing.T) {
	var (
		key           = "mykey"
		initialLength = 10
		updatedLength = 5
		r             = NewRecord(key, make([]byte, initialLength))
	)

	r.Update(make([]byte, updatedLength))

	if l := len(r.Value()); l != updatedLength {
		t.Errorf("Update failed. Expected length [%d], found [%d]\n", updatedLength, len(r.Value()))
	}
}

func TestNewRecordFrom(t *testing.T) {
	var (
		key           = "mykey"
		initialLength = 10
		updatedLength = 5
		r             = NewRecord(key, make([]byte, initialLength))
	)

	if l := len(r.Value()); l != initialLength {
		t.Errorf("Expected [%d]. Found [%d]\n", initialLength, l)
	}

	r2 := NewRecordFrom(r, make([]byte, updatedLength))

	if l := len(r2.Value()); l != updatedLength {
		t.Errorf("Expected [%d]. Found [%d]\n", updatedLength, l)
	}
}
