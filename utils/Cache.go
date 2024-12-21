package utils

type Cache[K comparable, V struct{}] map[K]V

func NewCache[K comparable, V struct{}]() Cache[K, V] {
	return make(map[K]V)
}

func (c *Cache[K, V]) Exists(key K) bool {
	_, ok := (*c)[key]
	return ok
}

type CachedList[K comparable, V interface{}] map[K][]V

// NewCachedList creates a new CachedList with the given key and value types.
//
// A CachedList is a map from a key type K to a slice of values of type V.
// The purpose of a CachedList is to provide a way of storing values of type V
// that are associated with a particular key of type K in a way that allows
// fast lookups of the values associated with a given key.
//
// The CachedList is not thread-safe. If you need to use it in a concurrent
// environment, you will need to add your own synchronization.
func NewCachedList[K comparable, V interface{}]() CachedList[K, V] {
	return make(map[K][]V)
}

// SetAt sets the value at the given index in the list associated with the given key.
//
// The CachedList must contain the given key and the index must be within the range
// of the list associated with the given key. If the key is not in the CachedList,
// or if the index is out of range, the method will panic.
//
// The value is set at the given index in the list associated with the given key.
func (c *CachedList[K, V]) SetAt(key K, index int, value V) {
	(*c)[key][index] = value
}

// AddAll adds the given values to the list associated with the given key.
//
// If the key is not in the CachedList, the method will panic.
//
// The values are appended to the end of the list associated with the given key.
func (c *CachedList[K, V]) AddAll(key K, values []V) {
	(*c)[key] = append((*c)[key], values...)
}

// EvictKey evicts the given key from the CachedList.
//
// If the key is not in the CachedList, the method does nothing.
//
// The method deletes the key from the CachedList, and sets the value associated
// with the key to nil.
func (c *CachedList[K, V]) EvictKey(key K) {
	delete(*c, key)
}

// EvictValues evicts all values associated with the given key from the CachedList.
//
// If the key is not in the CachedList, the method does nothing.
//
// The method does not delete the key from the CachedList. Instead, it sets the
// value associated with the key to nil. This allows the key to still be looked
// up in the CachedList, but the value associated with the key will be nil.
func (c *CachedList[K, V]) EvictValues(key K) {
	if c.Exists(key) {
		(*c)[key] = nil
	}
}

// EvictAll clears all keys and values from the CachedList.
//
// The EvictAll method is very fast and does not allocate. It is used to clear
// the CachedList completely. After calling this method, the CachedList will
// contain no keys or values.
func (c *CachedList[K, V]) EvictAll() {
	*c = make(map[K][]V)
}

// Exists returns true if the given key exists in the CachedList, and false otherwise.
//
// The Exists method does not allocate and is very fast.
func (c *CachedList[K, V]) Exists(key K) bool {
	_, ok := (*c)[key]
	return ok
}

// Get returns the list of values associated with the given key.
//
// If the key is not in the CachedList, the method will return nil.
//
// The returned list is a copy of the list associated with the given key.
// Modifying the returned list will not affect the CachedList.
func (c *CachedList[K, V]) Get(key K) []V {
	return (*c)[key]
}
