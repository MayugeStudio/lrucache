package lrucache

import (
	"container/list"
	// "fmt"
)

type Pair struct {
	First  string
	Second string
}

type LRUCache struct {
	l        *list.List
	capacity int
}

func New(capacity int) *LRUCache {
	lru := &LRUCache{
		capacity: capacity,
	}
	lru.l = list.New()
	return lru
}

func (lru *LRUCache) Put(first string, second string) {
	lru.l.PushFront(Pair{
		First:  first,
		Second: second,
	})

	if lru.l.Len() > lru.capacity {
		back := lru.l.Back()
		lru.l.Remove(back)
	}
}

func (lru *LRUCache) Get(key string) (string, bool) {
	for e := lru.l.Front(); e.Next() != nil; e = e.Next() {
		if e.Value.(Pair).First == key {
			lru.l.MoveToFront(e)
			return e.Value.(Pair).Second, true
		}
	}
	return "", false
}

// func (lru *LRUCache) Remove(key string) bool {
// 	lru.l.Remove(key)
// }
