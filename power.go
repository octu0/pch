package pch

import (
	"hash"
	"hash/fnv"
	"math"
	"math/bits"
	"unsafe"
)

type PowerConsistentHash struct {
	buckets      uint32
	minuesOne    uint32
	halfMinusOne uint32
	hasher       hash.Hash64
}

func (p *PowerConsistentHash) Hash(key string) uint32 {
	p.hasher.Reset()
	b := unsafe.Slice(unsafe.StringData(key), len(key))
	p.hasher.Write(b)
	h := p.hasher.Sum64()
	return p.get(h)
}

func (p *PowerConsistentHash) get(key uint64) uint32 {
	pool := newSimpleRandPool(key)
	r1 := f(key, p.minuesOne, pool)
	if r1 < p.buckets {
		return r1
	}
	r2 := g(key, p.buckets, p.halfMinusOne, pool)
	if p.halfMinusOne < r2 {
		return r2
	}
	return f(key, p.halfMinusOne, pool)
}

func New(n uint32, hasher hash.Hash64) *PowerConsistentHash {
	if n < 2 {
		panic("not enough buckets")
	}

	minuesOne := align16(n - 1)
	halfMinusOne := (n >> 1) - 1

	return &PowerConsistentHash{
		buckets:      n,
		minuesOne:    minuesOne,
		halfMinusOne: halfMinusOne,
		hasher:       hasher,
	}
}

func align16(m uint32) uint32 {
	m |= m >> 1
	m |= m >> 2
	m |= m >> 4
	m |= m >> 8
	m |= m >> 16
	return m
}

// k1 { kBits: 0001, j: 0, h: 1, R(k1, 0), range:[1] }
// k2 { kBits: 0010, j: 1, h: 2, R(k2, 1), range:[2,3] }
// k3 { kBits: 0001, j: 2, h: 4, R(k3, 2), range:[4,5,6,7] }
// k4 { kBits: 0001, j: 3, h: 8, R(k4, 3), range:[8,9,10,11,12,13,14,15] }
func f(key uint64, m uint32, pool generatorPool) uint32 {
	kBits := uint32(key & uint64(m))
	if kBits == 0 {
		return 0
	}

	j := uint32(findLastOne(kBits))
	h := uint32(1 << j)
	r := h + pool.Get().Rand(j)&(h-1)
	return r
}

// g returns [s, n - 1]
func g(key uint64, n uint32, s uint32, pool generatorPool) uint32 {
	x := s
	nn := uint64(n)
	for {
		k := (uint64(x) + 1) * uint64(math.MaxUint32)
		rndPlusOne := uint64(pool.Get().Next()) + 1

		if rndPlusOne*nn <= k {
			break
		}
		x = uint32(k / rndPlusOne)
	}
	return x
}

var (
	maxBitsUint32 int = bits.OnesCount32(uint32(math.MaxUint32))
)

func findLastOne(n uint32) int {
	return maxBitsUint32 - bits.LeadingZeros32(n)
}

var (
	FNV1  hash.Hash64 = fnv.New64()
	FNV1a hash.Hash64 = fnv.New64a()
)
