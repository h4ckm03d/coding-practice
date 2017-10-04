// Design and implement a data structure for Least Recently Used (LRU) cache. It should support the following operations: get and put.

// get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
// put(key, value) - Set or insert the value if the key is not already present. When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.

// Follow up:
// Could you do both operations in O(1) time complexity?

// Example:

// LRUCache cache = new LRUCache( 2 /* capacity */ );

// cache.put(1, 1);
// cache.put(2, 2);
// cache.get(1);       // returns 1
// cache.put(3, 3);    // evicts key 2
// cache.get(2);       // returns -1 (not found)
// cache.put(4, 4);    // evicts key 1
// cache.get(1);       // returns -1 (not found)
// cache.get(3);       // returns 3
// cache.get(4);       // returns 4

package lru
import (
	"container/list"
)

type LRUCache struct {
	size int
	items     map[interface{}]*list.Element
	evictList *list.List
}

// entry is used to hold a value in the evictList
type entry struct {
	key   int
	value int
}

func Constructor(size int) LRUCache {

	c := LRUCache{
		size:      size,
		evictList: list.New(),
		items:     make(map[interface{}]*list.Element),
	}
	return c
}


func (c *LRUCache) Get(key int) int {
	if ent, ok := c.items[key]; ok {
		c.evictList.MoveToFront(ent)
		return ent.Value.(*entry).value
	}
	return -1
}


func (c *LRUCache) Put(key int, value int)  {
	// Check for existing item
	if ent, ok := c.items[key]; ok {
		c.evictList.MoveToFront(ent)
		ent.Value.(*entry).value = value
		return
	}

	// Add new item
	ent := &entry{key, value}
	entry := c.evictList.PushFront(ent)
	c.items[key] = entry

	evict := c.evictList.Len() > c.size
	// Verify size not exceeded
	if evict {
		ent := c.evictList.Back()
		if ent != nil {
			c.removeElement(ent)
		}
	}
}

// removeElement is used to remove a given list element from the cache
func (c *LRUCache) removeElement(e *list.Element) {
	c.evictList.Remove(e)
	kv := e.Value.(*entry)
	delete(c.items, kv.key)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */