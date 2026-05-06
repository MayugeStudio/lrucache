package lrucache

import (
	"testing"
	"fmt"
	"strings"
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


func Benchmark_Put(b *testing.B) {
	lru := New(20)

	b.ResetTimer()

	for i := 0; i<b.N; i++ {
		lru.Put(fmt.Sprintf("Lower: %d", i), strings.ToUpper(fmt.Sprintf("Upper: %d", i)))
	}
}

func Benchmark_Get(b *testing.B) {
	lru := New(1000)
	for i := 0; i<1000; i++ {
		lru.Put(fmt.Sprintf("Lower: %d", i), strings.ToUpper(fmt.Sprintf("Upper: %d", i)))
	}

	b.ResetTimer()

	for i := 0; i<b.N; i++ {
		key := i % 1000
		lru.Get(fmt.Sprintf("Lower: %d", key))
	}
}

// This testing method is based on HashiCorp golang-lru benchmark 
func Benchmark_Random(b *testing.B) {
	lru := New(4096)
	trace := make([]string, b.N*2)
	for i := 0; i<b.N*2; i++ {
		trace[i] = fmt.Sprintf("%d", getRand(b) % 16384)
	}

	b.ResetTimer()

	var hit, miss int
	for i := 0; i<b.N*2; i++ {
		if i % 2 == 0 {
			lru.Put(trace[i], trace[i])
		} else {
			if _, ok := lru.Get(trace[i]); ok {
				hit += 1
			} else {
				miss += 1
			}
		}
	}

	b.Logf("hit: %d, miss: %d, ratio: %f", hit, miss, float64(hit)/float64(miss+hit))
}

