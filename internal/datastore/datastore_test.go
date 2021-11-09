package datastore

import (
	"errors"
	"testing"
)

func TestNewDatastore(t *testing.T) {

	// test: setup
	actual, err := NewDatastore()
	if err != nil || actual == nil {
		t.Fatalf("expected datastore to initialize without fault: %s", err)
	}

}

func TestDatastore_DescribeRecord(t *testing.T) {

	// test: setup
	ds, err := NewDatastore()
	if err != nil {
		t.Fatalf("expected datastore to initialize without fault: %s", err)
	}

	// test: ensure silly items aren't found
	_, err = ds.DescribeRecord("NOT_FOUND", Magic)
	if !errors.Is(err, ErrItemNotFound) {
		t.Fatalf("expected ErrItemNotFound: %s", err)
	}

	// test: ensure we can actually find stuff
	expected := "Firaga"
	actual, err := ds.DescribeRecord(expected, Magic)
	if err != nil {
		t.Fatalf("expected success: %s", err)
	}
	if expected != actual.Name {
		t.Fatalf("expected %s: %s", expected, actual.Name)
	}

	// test: ensure real things of the wrong type aren't found
	_, err = ds.DescribeRecord(expected, LimitBreak)
	if !errors.Is(err, ErrItemNotFound) {
		t.Fatalf("expected ErrItemNotFound: %s", err)
	}

}
