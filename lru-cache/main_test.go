package main

import "testing"

func TestLRUCacheCapacity(t *testing.T) {

	capacity := 2
	lruCache := (LRUCache{}).New(capacity)

	lruCache.insert("key1", "value1")
	lruCache.insert("key2", "value2")
	lruCache.insert("key3", "value3")

	value, _ := lruCache.get("key1")
	if value != "" {
		t.Errorf("Key 1 should not exist")
	}
}

func TestLRUCacheRecentlyUse(t *testing.T) {

	capacity := 2
	lruCache := (LRUCache{}).New(capacity)

	lruCache.insert("key1", "value1")
	lruCache.insert("key2", "value2")

	lruCache.get("key1")

	lruCache.insert("key3", "value3")

	value, _ := lruCache.get("key1")
	valueKey2, _ := lruCache.get("key2")
	if value == "" {
		t.Errorf("Key 1 should exist")
	}

	if valueKey2 != "" {
		t.Errorf("Key 2 least recently used should remove when insert key3")
	}
}

func TestLRUCacheRefreshInsert(t *testing.T) {

	capacity := 2
	lruCache := (LRUCache{}).New(capacity)

	lruCache.insert("key1", "value1")
	lruCache.insert("key2", "value2")

	lruCache.insert("key1", "new_value1")

	lruCache.insert("key3", "value3")

	value, _ := lruCache.get("key1")
	valueKey2, _ := lruCache.get("key2")
	if value != "new_value1" {
		t.Errorf("Key 1 value should update to newest one")
	}

	if valueKey2 != "" {
		t.Errorf("Key 2 least recently used should remove when insert key3")
	}
}
