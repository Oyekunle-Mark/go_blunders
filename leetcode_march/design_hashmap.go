package main

import (
	"bytes"
	"container/list"
	"strconv"
)

const BucketSize = 1000000

type MapValue struct {
	key   int
	value int
}

type MyHashMap struct {
	buckets [BucketSize]*list.List
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{
		buckets: [BucketSize]*list.List{},
	}
}

func (m *MyHashMap) Hash(key int) uint {
	var buf bytes.Buffer
	buf.WriteString(strconv.FormatInt(int64(key), 10))

	h := djb2Hash(&buf)

	return h % uint(BucketSize)
}

/** value will always be non-negative. */
func (m *MyHashMap) Put(key int, value int) {
	mapIndex := m.Hash(key)
	bucket := m.buckets[mapIndex]

	for e := bucket.Front(); e != nil; e = e.Next() {
		if e.Value.(MapValue).key == key {
			e.Value = MapValue{key, value}
			return
		}
	}

	m.buckets[mapIndex].PushBack(MapValue{key, value})
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (m *MyHashMap) Get(key int) int {
	mapIndex := m.Hash(key)
	bucket := m.buckets[mapIndex]

	for e := bucket.Front(); e != nil; e = e.Next() {
		if e.Value.(MapValue).key == key {
			return e.Value.(MapValue).value
		}
	}

	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (m *MyHashMap) Remove(key int) {
	mapIndex := m.Hash(key)
	bucket := m.buckets[mapIndex]

	for e := bucket.Front(); e != nil; e = e.Next() {
		if e.Value.(MapValue).key == key {
			bucket.Remove(e)
		}
	}
}

func djb2Hash(buf *bytes.Buffer) uint {
	var h uint = 5381

	for _, r := range buf.Bytes() {
		h = (h << 5) + h + uint(r)
	}

	return h
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
