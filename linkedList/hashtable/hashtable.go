package main

// pair represents a key-value pair stored inside a bucket.
type Pair struct {
	Key   string
	Value string
}

// Bucker is a slice of Pairs
type Bucket []Pair

// HashTable contains a fixed-size array of buckets
type HashTable struct {
	buckets []Bucket
	size int
}

// NewHashtable creates a new hashtable with given size.
func NewHashtable(size int) *HashTable {
	return &HashTable{
		buckets: make([]Bucket, size),
		size: size,
	}
}

// hash function: converts a string key into an index
func (h *HashTable) hash(key string) int {
	sum := 0
	for _, ch := range key {
		sum += int(ch)
	}
	return sum % h.size
}

// Insert adds or updates a key-value pair
func (h *HashTable) Insert(key, value string) {
	index := h.hash(key)
	bucket := h.buckets[index]

	
	// check if key already exists → update value
	for i, pair := range bucket {
		if pair.Key == key {
			bucket[i].Value = value
			h.buckets[index] = bucket
			return
		}
	}

	
	// otherwise add new key-value pair
	h.buckets[index] = append(bucket, Pair{Key: key, Value: value})
}

// Search returns the value of a key.
func (h *HashTable) Search(key string) (string, bool) {
	index := h.hash(key)
	bucket := h.buckets[index]

	for _, pair := range bucket {
		if pair.Key == key {
			return pair.Value, true
		}
	}
	return "", false
}


// Delete removes a key-value pair.
func (h *HashTable) Delete(key string) {
	index := h.hash(key)
	bucket := h.buckets[index]

	for i, pair := range bucket {
		if pair.Key == key {
			h.buckets[index] = append(bucket[:i], bucket[i+1:]...)
			return
		}
	}
}
