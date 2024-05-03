package pch

import (
	"pgregory.net/rand"
)

type generator interface {
	Rand(uint32) uint32
	Next() uint32
}

var (
	_ generator = (*pchRand)(nil)
)

type pchRand struct {
	rnd *rand.Rand
}

func (p *pchRand) Rand(j uint32) uint32 {
	return p.rnd.Uint32n(j)
}

func (p *pchRand) Next() uint32 {
	return p.rnd.Uint32()
}

func newRand(key uint64) *pchRand {
	return &pchRand{rand.New(key)}
}

type generatorPool interface {
	Get() generator
}

var (
	_ generatorPool = (*simpleRandPool)(nil)
)

type simpleRandPool struct {
	key uint64
	rnd *pchRand
}

func (p *simpleRandPool) Get() generator {
	if p.rnd == nil {
		p.rnd = newRand(p.key)
	}
	return p.rnd
}

func newSimpleRandPool(key uint64) generatorPool {
	return &simpleRandPool{key, nil}
}
