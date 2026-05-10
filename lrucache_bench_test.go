package lrucache_test

import (
	"fmt"
	"strings"
	"testing"
	"github.com/MayugeStudio/lrucache"
)

func Benchmark_Put(b *testing.B) {
	lru := lrucache.New(20)

	b.ResetTimer()

	for i := 0; i<b.N; i++ {
		lru.Put(fmt.Sprintf("Lower: %d", i), strings.ToUpper(fmt.Sprintf("Upper: %d", i)))
	}
}

func Benchmark_Get(b *testing.B) {
	lru := lrucache.New(1000)
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
	lru := lrucache.New(4096)
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
