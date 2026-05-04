package lrucache

import (
	"testing"
)

func TestLRUCache_Put(t *testing.T) {
	lru := New()
	lru.Put("<abc>", "</abc>")
	lru.Put("<def>", "</def>")
	lru.Put("<ghi>", "</ghi>")

	if lru.l.Len() != 3 {
		t.Errorf("Expected 3 but got %d", lru.l.Len())
	}

	if lru.l.Front().Value.(Pair).First != "<abc>" {
		t.Errorf("Expected 10 but got %d", lru.l.Front().Value)
	}

	if lru.l.Front().Next().Value.(Pair).First != "<def>" {
		t.Errorf("Expected 20 but got %d", lru.l.Front().Next().Value)
	}

	if lru.l.Front().Next().Next().Value.(Pair).First != "<ghi>" {
		t.Errorf("Expected 30 but got %d", lru.l.Front().Next().Next().Value)
	}
}

func Test_LRUCache_Get(t *testing.T) {
	lru := New()
	lru.Put("<abc>", "</abc>")
	lru.Put("<def>", "</def>")
	lru.Put("<ghi>", "</ghi>")

	if lru.l.Len() != 3 {
		t.Errorf("Expected 3 but got %d", 3)
	}
	
	if r, ok := lru.Get("<abc>"); ok && r != "</abc>" {
		t.Errorf("expected </abc> but got %s", r)
	}
}

