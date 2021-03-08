package main

import (
	"bytes"
	"container/list"
	"strconv"
)

const BucketSize = 1000000

// MapPair represents the hash map key, value pair
type MapPair struct {
	key   int
	value int
}

// MyHashMap is the hash map type
type MyHashMap struct {
	buckets [BucketSize]*list.List
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{
		buckets: [BucketSize]*list.List{},
	}
}

// Hash returns the hash of the key
func (m *MyHashMap) Hash(key int) uint {
	var buf bytes.Buffer
	buf.WriteString(strconv.FormatInt(int64(key), 10))

	h := djb2Hash(&buf)

	return h % uint(BucketSize) // return hash modulo bucket size to limit to bucket size
}

/** value will always be non-negative. */
func (m *MyHashMap) Put(key int, value int) {
	mapIndex := m.Hash(key)
	bucket := m.buckets[mapIndex]

	// if current bucket hasn't been initialized with a dll
	if bucket == nil {
		m.buckets[mapIndex] = list.New()                  // initialize
		m.buckets[mapIndex].PushBack(MapPair{key, value}) // insert key, value pair
	} else {
		// search all dll elements for the key
		for e := bucket.Front(); e != nil; e = e.Next() {
			// if current element has the key
			if e.Value.(MapPair).key == key {
				e.Value = MapPair{key, value} // update and return with new pair
				return
			}
		}

		// no element found in dll with the key
		m.buckets[mapIndex].PushBack(MapPair{key, value}) // insert and return
		return
	}
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (m *MyHashMap) Get(key int) int {
	mapIndex := m.Hash(key)
	bucket := m.buckets[mapIndex]

	// return early if bucket is uninitialized
	if bucket == nil {
		return -1
	}

	// loop through dll elements for key
	for e := bucket.Front(); e != nil; e = e.Next() {
		// if found, return it's value
		if e.Value.(MapPair).key == key {
			return e.Value.(MapPair).value
		}
	}

	return -1 // not found in dll at the bucket
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (m *MyHashMap) Remove(key int) {
	mapIndex := m.Hash(key)
	bucket := m.buckets[mapIndex]

	// return early if bucket is uninitialized
	if bucket == nil {
		return
	}

	// loop through dll elements for key
	for e := bucket.Front(); e != nil; e = e.Next() {
		// call remove on dll with element if it has the key
		if e.Value.(MapPair).key == key {
			bucket.Remove(e)
		}
	}

}

// djb2Hash creates the hash
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
