package lrucache

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	l        *list.List // linked list of entries
	cache    map[string]*list.Element
	capacity int
}

type entry struct {
	key   string
	value string
}

func New(capacity int) *LRUCache {
	lru := &LRUCache{
		l: list.New(),
		cache: make(map[string]*list.Element, capacity),
		capacity: capacity,
	}
	return lru
}

func (lru *LRUCache) Put(key string, value string) {
	if element, ok := lru.cache[key]; ok {
		// Update an entry value with new one and move it to the front of the list
		element.Value.(*entry).value = value
		lru.l.MoveToFront(element)
		return
	}

	// Remove the least recently used element if the cache size exceeds its capacity.
	if lru.l.Len() == lru.capacity {
		// TODO: Handle back == nil situation
		back := lru.l.Back()
		delete(lru.cache, back.Value.(*entry).key)
		lru.l.Remove(back)
	}

	// Add an element if the key is not in the cache.
	e := lru.l.PushFront(&entry{ key: key, value: value })
	lru.cache[key] = e
}

func (lru *LRUCache) Get(key string) (string, bool) {
	if element, ok := lru.cache[key]; ok {
		lru.l.MoveToFront(element)
		return element.Value.(*entry).value, true
	}

	return "", false
}

// Count gives the current count of the elements in this cache
func (lru *LRUCache) Count() int {
	return lru.l.Len()
}

// Dump prints the elements in the cache.
func (lru *LRUCache) Dump() {
	for e := lru.l.Front(); e != nil; e = e.Next() {
		fmt.Println("%s => %s\n", e.Value.(entry).key, e.Value.(entry).value)
	}
}

