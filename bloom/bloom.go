package bloom

import (
	"hash"
	"hash/fnv"
	"math"
)

// BloomFilter
// https://en.wikipedia.org/wiki/Bloom_filter
type BloomFilter struct {
	h hash.Hash32

	k int   // k is number of hash functions
	m int64 // m is filter size in bits

	bucket int64
	states []uint32
}

// New creates a new Bloom Filter. n is number of elements
// you want to store and p is false-positive probability.
//
// When you store data on m bits array, the probability of
// collision is 1/m, so not-collision is 1 - (1/m).
//
// When you use k hash functions and store n elements,
// the probalibity of not collision is:
//
//    (1 - (1/m))^(k*n)
//
// In this case, when you inject non-mebers elements,
// the probalitity of all bit becames 1 is:
//
//    (1 - (1 - (1/m))^(k*n) )^k = (1 - e^(-kn/m))^k
//
// For a given m and n, the value of the number of hash functions
// k that minimizes the false positive probability is,
//
//    k = (m/n)ln(2)
//
// The required number of bits m, the number of inserted elements
// and a desired false positive probability p (and assuming the optimal
// value of k is used) can be)
//
//    m =  n * log2(e) * log2(1/p)
//
func New(n, p int) *BloomFilter {
	// TODO(tcnksm): error handling
	m := int64(math.Abs(math.Ceil(float64(n) * math.Log2(math.E) * math.Log2(1/float64(p)))))
	k := int(float64(m/int64(n)) * math.Log2(2))
	return newFilter(k, m)
}

// New creates Bloom Filter with m bits and k hashing functions.
func newFilter(k int, m int64) *BloomFilter {
	bucket := m / 32
	return &BloomFilter{
		h:      fnv.New32a(),
		k:      k,
		m:      m,
		bucket: bucket,
		states: make([]uint32, uint(bucket)),
	}
}

// Add adds data to the filter.
func (f *BloomFilter) Add(data []byte) error {
	for i := 1; i < (f.k + 1); i++ {
		f.h.Write(data)
		v := f.h.Sum32()
		f.h.Reset()

		index := int64(v*uint32(i)) % f.m // [0,m]
		bucket := (index / 32) % f.bucket // [0,bucket]
		offset := index % 32              // [0,32]
		f.states[bucket] = f.states[bucket] | 1<<uint(offset)
	}
	return nil
}

// Contains returns the given data is added or not.
func (f *BloomFilter) Contains(data []byte) bool {
	for i := 1; i < (f.k + 1); i++ {
		f.h.Write(data)
		v := f.h.Sum32()
		f.h.Reset()

		index := int64(v*uint32(i)) % f.m // [0,m]
		bucket := (index / 32) % f.bucket // [0,bucket]
		offset := index % 32              // [0,32]
		if f.states[bucket]&(1<<uint(offset)) == 0 {
			return false
		}
	}
	return true
}
