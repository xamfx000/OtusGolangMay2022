package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	mx       sync.Mutex
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mx.Lock()
	defer c.mx.Unlock()
	newCacheItem := cacheItem{key: key, value: value}
	item, ok := c.items[key]
	if ok {
		item.Value = newCacheItem
		c.queue.MoveToFront(item)
		return true
	}
	if c.queue.Len() >= c.capacity {
		itemToDelete := c.queue.Back().Value.(cacheItem)
		delete(c.items, itemToDelete.key)
		c.queue.Remove(c.queue.Back())
	}
	newCacheListItem := &ListItem{Value: newCacheItem}
	c.items[key] = newCacheListItem
	c.queue.PushFront(newCacheListItem.Value)
	return ok
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mx.Lock()
	defer c.mx.Unlock()
	item, ok := c.items[key]
	if ok {
		c.queue.MoveToFront(item)
		cacheHitItem := item.Value.(cacheItem)
		return cacheHitItem.value, ok
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.items = map[Key]*ListItem{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
