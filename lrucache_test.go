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

	if lru.l.Front().Value.(*entry).key != "c" {
		t.Fatalf("Expected `c` but got %s\n", lru.l.Front().Value.(*entry).key)
	}

	if element, ok := lru.cache["c"]; !ok {
		t.Fatalf("Expected `c` to be in the cache but it is not\n")
	} else if element.Value.(*entry).key != "c" || 
						element.Value.(*entry).value != "C" {
		t.Fatalf("Expected `C` but got %s\n", element.Value.(*entry).value)
	}


	if lru.l.Front().Next().Value.(*entry).key != "b" {
		t.Fatalf("Expected `b` but got %s\n", lru.l.Front().Next().Value.(*entry).key)
	}

	if element, ok := lru.cache["b"]; !ok {
		t.Fatalf("Expected `b` to be in the cache but it is not\n")
	} else if element.Value.(*entry).key != "b" ||
						element.Value.(*entry).value != "B" {
		t.Fatalf("Expected `B` but got %s\n", element.Value.(*entry).value)
	}


	if lru.l.Front().Next().Next().Value.(*entry).key != "a" {
		t.Fatalf("Expected `a` but got %s\n", lru.l.Front().Next().Next().Value.(*entry).key)
	}

	if element, ok := lru.cache["a"]; !ok {
		t.Fatalf("Expected `a` to be in the cache but it is not\n")
	} else if element.Value.(*entry).key != "a" {
		t.Fatalf("Expected `a` but got %s\n", element.Value.(*entry).value)
	} else if element.Value.(*entry).value != "A" {
		t.Fatalf("Expected `A` but got %s\n", element.Value.(*entry).value)
	}
}

func Test_LRUCache_Put_RemoveLastElementIfLenExceedCapacity(t *testing.T) {
	lru := New(5)
	lru.Put("a", "A") // l: a
	lru.Put("b", "B") // l: b a
	lru.Put("c", "C") // l: c b a
	lru.Put("d", "D") // l: d c b a
	lru.Put("e", "E") // l: e d c b a
	if lru.l.Len() != 5 {
		t.Fatalf("Expected 5 but got %d\n", lru.l.Len())
	}

	lru.Put("f", "F") // expected to be like this: f e d c b

	if lru.l.Len() != 5 {
		t.Fatalf("Expected 5 but got %d\n", lru.l.Len())
	}

	keys := []string{"f", "e", "d", "c", "b"}
	values := []string{"F", "E", "D", "C", "B"}

	e := lru.l.Front()
	for i := range 5 {
		if e == nil {
			t.Fatalf("Expected something but got nil\n")
		}

		key := keys[i]
		expectedValue := values[i]
		if element, ok := lru.cache[key]; !ok {
			t.Fatalf("Expected `%s` to be in the cache but it is not", key)
		} else if element.Value.(*entry).value != expectedValue {
			t.Fatalf("Expected `%s` but got %s\n", element.Value.(*entry).value, expectedValue)
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
		t.Fatalf("Expected A but got %s\n", r)
	}
}

func Test_LRUCache_Get_nil(t *testing.T) {
	lru := New(10)
	if _, ok := lru.Get("a"); ok {
		t.Fatalf("Expected not to be ok but got ok")
	}
}

func Test_LRUCache_Cap1(t *testing.T) {
	lru := New(1)
	lru.Put("a", "A")
	if lru.l.Len() != 1 {
		t.Fatalf("Expected 1 but got %d\n", lru.l.Len())
	}

	if value, ok := lru.Get("a"); !ok {
		t.Fatalf("Expected ok but got not ok\n")
	} else if value != "A" {
		t.Fatalf("Expected `A` but got %s\n", value)
	}

	lru.Put("b", "B")
	
	if _, ok := lru.Get("a"); ok {
		t.Fatalf("Expected not ok but got ok\n")
	}

	if value, ok := lru.Get("b"); !ok {
		t.Fatalf("Expected ok but got not ok\n")
	} else if value != "B" {
		t.Fatalf("Expected `B` but got %s\n", value)
	}
}

func Test_LRUCache_Size(t *testing.T) {
	lru := New(10)
	lru.Put("a", "A")
	lru.Put("b", "B")
	lru.Put("c", "C")

	if lru.Count() != lru.l.Len() {
		t.Fatalf("Expected %d but got %d\n", lru.l.Len(), lru.Count())
	}
}

func Test_LRUCache_Dump(t *testing.T) {
	lru := New(10)
	lru.Put("a", "A")
	lru.Put("b", "B")
	lru.Put("c", "C")

	lru.Dump()
}

