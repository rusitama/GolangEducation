package main

import (
	"fmt"
)

type Cache interface {
	Get(k string) (string, bool)
	Set(k, v string)
}

var _ Cache = (*cacheImpl)(nil)

func newCacheImpl() *cacheImpl {
	return &cacheImpl{
		data: make(map[string]string),
	}
}

type cacheImpl struct {
	data map[string]string
}

func (c *cacheImpl) Get(k string) (string, bool) {
	v, ok := c.data[k]
	return v, ok
}

func (c *cacheImpl) Set(k, v string) {
	c.data[k] = v
}

func newDbImpl(cache Cache) *dbImpl {
	return &dbImpl{
		cache: cache,
		dbs:   map[string]string{"hello": "world", "test": "test"},
	}
}

type dbImpl struct {
	cache Cache
	dbs   map[string]string
}

func (d *dbImpl) Get(k string) (string, bool) {
	v, ok := d.cache.Get(k)
	if ok {
		return fmt.Sprintf("answer from cache: key: %s, val: %s", k, v), ok
	}

	v, ok = d.dbs[k]
	return fmt.Sprintf("answer from dbs: key: %s, val: %s", k, v), ok
}

func main() {
	c := newCacheImpl()
	db := newDbImpl(c)

	// Fill cache with data from db
	for k, v := range db.dbs {
		c.Set(k, v)
	}

	// Test cache functionality
	fmt.Println(db.Get("test"))
	fmt.Println(db.Get("hello"))
}
