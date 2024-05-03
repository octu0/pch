# `pch`

[![MIT License](https://img.shields.io/github/license/octu0/pch)](https://github.com/octu0/pch/blob/master/LICENSE)
[![GoDoc](https://pkg.go.dev/badge/github.com/octu0/pch)](https://pkg.go.dev/github.com/octu0/pch)
[![Go Report Card](https://goreportcard.com/badge/github.com/octu0/pch)](https://goreportcard.com/report/github.com/octu0/pch)
[![Releases](https://img.shields.io/github/v/release/octu0/pch)](https://github.com/octu0/pch/releases)

Go implementation of [Power Consistent Hashing](https://arxiv.org/pdf/2307.12448.pdf) algorithm.

## Installation

```bash
go get github.com/octu0/pch
```

## Example

```go
import (
    "hash/fnv"

    "github.com/octu0/pch"
)

func main() {
    p := pch.New(512, fnv.New64()) // 512 buckets, hash function fnv.New64()
    p.Hash("hello world")
}
```

## Benchmark

```
$ go test -bench=Benchmark .
goos: linux
goarch: amd64
pkg: github.com/octu0/pch
cpu: Intel(R) Xeon(R) W-11955M CPU @ 2.60GHz
Benchmark/jump/512-16           15240927                76.51 ns/op
Benchmark/jump/1024-16          14898984                79.36 ns/op
Benchmark/jump/2048-16          14624468                81.27 ns/op
Benchmark/jump/4096-16          14171553                84.87 ns/op
Benchmark/jump/8192-16          13739763                86.91 ns/op
Benchmark/power/512-16           9934183               117.4 ns/op
Benchmark/power/1024-16         10374721               119.1 ns/op
Benchmark/power/2048-16         10252898               117.2 ns/op
Benchmark/power/4096-16         10489959               118.0 ns/op
Benchmark/power/8192-16         10361293               117.0 ns/op
PASS
```

# License

MIT, see LICENSE file for details.