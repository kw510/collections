package treeset_test

import (
	"testing"

	"github.com/kw510/collections/treeset"
)

func TestInsertValue(t *testing.T) {
	ts := treeset.New[string]()
	if !ts.InsertValue("test") {
		t.Errorf("should insert string")
	}
	if ts.InsertValue("test") {
		t.Errorf("should not insert string")
	}
	if ts.Len() != 1 {
		t.Errorf("Length should be 1")
	}

	tsv := treeset.New[string]()
	if !ts.InsertNode(tsv) {
		t.Errorf("should insert treeset")
	}
	if ts.InsertNode(tsv) {
		t.Errorf("should not insert treeset")
	}
	if ts.Len() != 2 {
		t.Errorf("Length should be 2")
	}
}

func TestContains(t *testing.T) {
	ts := treeset.New[any]()
	if !ts.InsertValue("test") {
		t.Errorf("should insert string")
	}
	if !ts.Contains("test") {
		t.Errorf("should contain string")
	}

	tsv := treeset.New[any]()
	if !ts.InsertNode(tsv) {
		t.Errorf("should insert treeset")
	}
	if !ts.Contains(tsv) {
		t.Errorf("should contain treeset")
	}
}

func TestRemove(t *testing.T) {
	ts := treeset.New[any]()
	ts.InsertValue("test")
	if !ts.Remove("test") {
		t.Errorf("should remove string")
	}
	if ts.Remove("test") {
		t.Errorf("should not remove string")
	}

	tsv := treeset.New[any]()
	ts.InsertNode(tsv)
	if !ts.Remove(tsv) {
		t.Errorf("should remove treeset")
	}
	if ts.Remove(tsv) {
		t.Errorf("should not remove string")
	}
}

func TestIsSubset(t *testing.T) {
	ts := treeset.New[any]()
	tsv := treeset.New[any]()

	ts.InsertNode(tsv)
	if !ts.IsSubset(tsv) {
		t.Errorf("tsv should be a subset of ts")
	}
	if !ts.IsSubset(ts) {
		t.Errorf("ts should be a subset of ts")
	}
	if !tsv.IsSubset(tsv) {
		t.Errorf("tsv should be a subset of tsv")
	}
	if tsv.IsSubset(ts) {
		t.Errorf("ts should not be a subset of tsv")
	}
}

func TestIsSuperSet(t *testing.T) {
	ts := treeset.New[any]()
	tsv := treeset.New[any]()

	ts.InsertNode(tsv)
	if !tsv.IsSuperSet(ts) {
		t.Errorf("ts should be a superset of tsv")
	}
	if !ts.IsSuperSet(ts) {
		t.Errorf("ts should be a superset of ts")
	}
	if !tsv.IsSuperSet(tsv) {
		t.Errorf("tsv should be a superset of tsv")
	}
	if ts.IsSuperSet(tsv) {
		t.Errorf("tsv should not be a superset of ts")
	}
}
