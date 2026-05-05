// TODO: Add newline character to the error messages.
// TODO: Use Fatalf instead of Errorf at the some places.
package lrucache

import (
	"testing"
)

func TestLRUCache_Put(t *testing.T) {
	lru := New(10)
	lru.Put("a", "A")
	lru.Put("b", "B")
	lru.Put("c", "C")

	if lru.l.Len() != 3 {
		t.Errorf("Expected 3 but got %d", lru.l.Len())
	}

	if lru.l.Front().Value.(Pair).First != "c" {
		t.Errorf("Expected `c` but got %s", lru.l.Front().Value.(Pair).First)
	}

	if lru.l.Front().Next().Value.(Pair).First != "b" {
		t.Errorf("Expected `b` but got %s", lru.l.Front().Next().Value.(Pair).First)
	}

	if lru.l.Front().Next().Next().Value.(Pair).First != "a" {
		t.Errorf("Expected `a` but got %s", lru.l.Front().Next().Next().Value.(Pair).First)
	}
}

func Test_LRUCache_Get(t *testing.T) {
	lru := New(10)
	lru.Put("a", "A")
	lru.Put("b", "B")
	lru.Put("c", "C")

	if lru.l.Len() != 3 {
		t.Errorf("Expected 3 but got %d", 3)
	}
	
	if r, ok := lru.Get("a"); ok && r != "A" {
		t.Errorf("expected A but got %s", r)
	}
}

