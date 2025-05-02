package oldcache

import (
	"fmt"
	"log/slog"
	"strings"
)

type CacheItem[T any] struct {
	Value T
	Key   int
	Next  *CacheItem[T]
	Prev  *CacheItem[T]
}

type Cache[T any] struct {
	Base []CacheItem[T]
	Head *CacheItem[T]
	Tail *CacheItem[T]
	Dict map[int]*CacheItem[T]
}

func (ll *Cache[T]) String() string {
	acc := []string{}
	for item := ll.Head; item != nil; item = item.Next {
		acc = append(acc, fmt.Sprintf("%v - %v", item.Key, item.Value))
	}

	return strings.Join(acc, ", ")
}

func NewCache[T any](size int) *Cache[T] {
	if size < 2 {
		panic("min size for linked list is 2")
	}

	c := &Cache[T]{}

	base := make([]CacheItem[T], size)
	c.Base = base

	c.Head = &base[0]
	c.Tail = &base[size-1]

	for i := range size {
		base[i].Key = -1

		if i < size-1 {
			base[i].Next = &base[i+1]
		}

		if i > 0 {
			base[i].Prev = &base[i-1]
		}

		slog.Debug("linked list item set", slog.Any("item", base[i]), slog.Int("i", i))
	}

	slog.Debug("linked list created", slog.String("list", c.String()))

	c.Dict = make(map[int]*CacheItem[T])

	return c
}

func (c *Cache[T]) Put(key int, value T) {
	if key < 0 {
		panic("cache key must be positive")
	}

	slog.Debug("started linked list update", slog.String("list", c.String()))

	newHead := c.Tail
	newTail := newHead.Prev

	slog.Debug("started tail update update", slog.Any("newTail", newTail))

	if newHead.Key != -1 {
		delete(c.Dict, newHead.Key)
	}

	newTail.Next = nil
	c.Tail = newTail

	newHead.Next = c.Head
	c.Head.Prev = newHead

	c.Head = newHead

	newHead.Value = value
	newHead.Key = key

	c.Dict[key] = c.Head

	slog.Debug("end update cache", slog.Any("dict", c.Dict))
}

func (c *Cache[T]) Get(key int) (*T, bool) {
	item, ok := c.Dict[key]
	var v *T
	if item != nil {
		v = &item.Value
	}

	if ok {
		prev := item.Prev
		next := item.Next

		if prev != nil {
			prev.Next = next
		}
		if next != nil {
			next.Prev = prev
		}

		item.Next = c.Head
		c.Head.Prev = item
		c.Head = item
	}

	slog.Debug("end cache search", slog.Any("dict", c.Dict))

	return v, ok
}
