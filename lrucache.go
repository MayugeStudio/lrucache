package lrucache

import (
	"container/list"
	// "fmt"
)

type Pair struct {
	First string
	Second string
}

type LRUCache struct {
	l *list.List
}

func New() *LRUCache {
	lru := &LRUCache{}
	lru.l = list.New()
	return lru
}

func (lru *LRUCache) Put(first string, second string) {
	lru.l.PushBack(Pair{
		First: first,
		Second: second,
	})
}

func (lru *LRUCache) Get(key string) (string, bool) {
	for e := lru.l.Front(); e.Next() != nil; e = e.Next() {
		if e.Value.(Pair).First == key {
			return e.Value.(Pair).Second, true
		}
	}
	return "", false
}

// func (lru *LRUCache) Remove(key string) bool {
// 	lru.l.Remove(key)
// }

