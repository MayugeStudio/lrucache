package main

import (
	"container/list"
	// "fmt"
)

type LRUCache struct {
	l *list.List
}

func New() *LRUCache {
	lru := &LRUCache{}
	lru.l = list.New()
	return lru
}

func (lru *LRUCache) Put(e any) {
	lru.l.PushBack(e)
}

// func (l LRUCache) Get() {}

func main() {
	
}

