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
		t.Fatalf("Expected 3 but got %d\n", lru.l.Len())
	}

	if lru.l.Front().Value.(Pair).First != "c" {
		t.Fatalf("Expected `c` but got %s\n", lru.l.Front().Value.(Pair).First)
	}

	if lru.l.Front().Next().Value.(Pair).First != "b" {
		t.Fatalf("Expected `b` but got %s\n", lru.l.Front().Next().Value.(Pair).First)
	}

	if lru.l.Front().Next().Next().Value.(Pair).First != "a" {
		t.Fatalf("Expected `a` but got %s\n", lru.l.Front().Next().Next().Value.(Pair).First)
	}
}

func Test_LRUCache_Put_RemoveLastElementIfLenExceedCapacity(t *testing.T) {
	lru := New(5)
	lru.Put("a", "A") // a
	lru.Put("b", "B") // b a
	lru.Put("c", "C") // c b a
	lru.Put("d", "D") // d c b a
	lru.Put("e", "E") // e d c b a
	if lru.l.Len() != 5 {
		t.Fatalf("Expected 5 but got %d\n", lru.l.Len())
	}

	lru.Put("f", "F") // expected to be like this: f e d c b

	if lru.l.Len() != 5 {
		t.Fatalf("Expected 5 but got %d\n", lru.l.Len())
	}

	expectedFirst := []string{"f", "e", "d", "c", "b"}
	expectedSecond := []string{"F", "E", "D", "C", "B"}

	e := lru.l.Front()
	for i := range 5 {
		if e == nil {
			t.Fatalf("Expected something but got nil\n")
		}

		first := expectedFirst[i]
		second := expectedSecond[i]
		if e.Value.(Pair).First != first {
			t.Fatalf("Expected `%s` but got %s\n", first, e.Value.(Pair).First)
		}
		if e.Value.(Pair).Second != second {
			t.Fatalf("Expected `%s` but got %s\n", second, e.Value.(Pair).Second)
		}

		e = e.Next()
	}
	
}

func Test_LRUCache_Get(t *testing.T) {
	lru := New(10)
	lru.Put("a", "A")
	lru.Put("b", "B")
	lru.Put("c", "C")

	if lru.l.Len() != 3 {
		t.Fatalf("Expected 3 but got %d\n", 3)
	}
	
	if r, ok := lru.Get("a"); ok && r != "A" {
		t.Fatalf("expected A but got %s\n", r)
	}
}

