package pch

import (
	"hash/fnv"
	"strconv"
	"testing"

	"github.com/lithammer/go-jump-consistent-hash"
)

func Benchmark(b *testing.B) {
	benchJump := func(b *testing.B, n int) {
		j := jump.New(n, jump.FNV1)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = j.Hash(strconv.Itoa(i))
		}
	}
	benchPower := func(b *testing.B, n uint32) {
		p := New(n, fnv.New64())
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = p.Hash(strconv.Itoa(i))
		}
	}
	b.Run("jump/512", func(tb *testing.B) {
		benchJump(tb, 512)
	})
	b.Run("jump/1024", func(tb *testing.B) {
		benchJump(tb, 1024)
	})
	b.Run("jump/2048", func(tb *testing.B) {
		benchJump(tb, 2048)
	})
	b.Run("jump/4096", func(tb *testing.B) {
		benchJump(tb, 4096)
	})
	b.Run("jump/8192", func(tb *testing.B) {
		benchJump(tb, 8192)
	})
	b.Run("power/512", func(tb *testing.B) {
		benchPower(tb, 512)
	})
	b.Run("power/1024", func(tb *testing.B) {
		benchPower(tb, 1024)
	})
	b.Run("power/2048", func(tb *testing.B) {
		benchPower(tb, 2048)
	})
	b.Run("power/4096", func(tb *testing.B) {
		benchPower(tb, 4096)
	})
	b.Run("power/8192", func(tb *testing.B) {
		benchPower(tb, 8192)
	})
}

func TestNew(t *testing.T) {
	t.Run("consistency", func(tt *testing.T) {
		a := New(3, fnv.New64())
		b := New(3, fnv.New64())

		v1 := a.Hash("test")
		v2 := b.Hash("test")
		v3 := a.Hash("test")
		if v1 != v2 {
			tt.Errorf("expected %v, got %v", v1, v2)
		}
		if v1 != v3 {
			tt.Errorf("expected %v, got %v", v1, v3)
		}

		v4 := a.Hash("test-2") //fnv64
		if v1 == v4 {
			tt.Errorf("expected %v, got %v", v1, v4)
		}
	})
	t.Run("no-consistency", func(tt *testing.T) {
		a := New(3, fnv.New64())
		b := New(30, fnv.New64())

		v1 := a.Hash("test")
		v2 := b.Hash("test")
		if v1 == v2 {
			tt.Errorf("expected %v, got %v", v1, v2)
		}
	})
}
