package main

import (
	"testing"
)

func TestLRUCachePut(t *testing.T) {
	lru := New()
	lru.Put(10)
	lru.Put(20)
	lru.Put(30)

	if (lru.l.Len() != 3) {
		t.Errorf("Expected 3 but got %d", lru.l.Len())
	}

	if lru.l.Front().Value != 10 {
		t.Errorf("Expected 10 but got %d", lru.l.Front().Value)
	}

	if lru.l.Front().Next().Value != 20 {
		t.Errorf("Expected 20 but got %d", lru.l.Front().Next().Value)
	}

	if lru.l.Front().Next().Next().Value != 30 {
		t.Errorf("Expected 30 but got %d", lru.l.Front().Next().Next().Value)
	}
}

