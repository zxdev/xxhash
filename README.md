# xxhash

The ```xxhash``` package implements the 64-bit variant of xxHash (XXH64) as described at http://cyan4973.github.io/xxHash/ without any Go standard library package dependancy nor any hash.Hash interface components.

```golang
// signatures
Sum(b []byte) uint64
SSum(s string) uint64
```


