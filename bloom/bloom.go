package bloom

import (
	"hash"
	"hash/fnv"
)

// BloomFilter
// https://en.wikipedia.org/wiki/Bloom_filter
type BloomFilter struct {
	h hash.Hash64

	k uint // k is number of hash functions
	m uint // m is filter size in bits

	bucket uint
	states []uint64
}

// New creates Bloom Filter with m bits and k hashing functions.
func New(k, m uint) *BloomFilter {
	bucket := m / 64
	return &BloomFilter{
		h:      fnv.New64a(),
		k:      k,
		m:      m,
		bucket: bucket,
		states: make([]uint64, bucket),
	}
}

// Add adds data to the filter.
func (f *BloomFilter) Add(data []byte) error {
	for i := uint(0); i < f.k; i++ {
		f.h.Write(data)
		v := f.h.Sum64()

		index := (uint(v) * i) % f.m      // [0,m]
		bucket := (index / 64) % f.bucket // [0,bucket]
		offset := index % 64              // [0,64]
		f.states[bucket] = f.states[bucket] | 1<<uint(offset)
		f.h.Reset()
	}
	return nil
}

// Contains returns the given data is added or not.
func (f *BloomFilter) Contains(data []byte) bool {
	for i := uint(0); i < f.k; i++ {
		f.h.Write(data)
		v := f.h.Sum64()

		index := (uint(v) * i) % f.m      // [0,m]
		bucket := (index / 64) % f.bucket // [0,bucket]
		offset := index % 64              // [0,64]
		if f.states[bucket]&(1<<uint(offset)) == 0 {
			return false
		}
		f.h.Reset()
	}
	return true
}
