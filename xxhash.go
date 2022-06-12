package xxhash

// xxhash implements the 64-bit variant of xxHash (XXH64)
// as described at http://cyan4973.github.io/xxHash/ without
// any standard library package dependancy

// SSum computes the uint64 xxHash digest of s
func SSum(s string) uint64 { return Sum([]byte(s)) }

// Sum computes the uint64 xxHash digest of b
func Sum(b []byte) uint64 {

	var n, h uint64 = uint64(len(b)), 2870177450012600261

	if n >= 32 {
		var v1 uint64 = 6983438078262162902
		var v2 uint64 = 14029467366897019727
		var v3 uint64 = 0
		var v4 uint64 = 7046029288634856825
		for len(b) >= 32 {
			v1 = round(v1, u64(b[0:8:len(b)]))
			v2 = round(v2, u64(b[8:16:len(b)]))
			v3 = round(v3, u64(b[16:24:len(b)]))
			v4 = round(v4, u64(b[24:32:len(b)]))
			b = b[32:len(b):len(b)]
		}
		h = (v1<<1 | v1>>63) + (v2<<7 | v2>>57) + (v3<<12 | v3>>52) + (v4<<18 | v4>>46)
		h = (h^round(0, v1))*11400714785074694791 + 9650029242287828579
		h = (h^round(0, v2))*11400714785074694791 + 9650029242287828579
		h = (h^round(0, v3))*11400714785074694791 + 9650029242287828579
		h = (h^round(0, v4))*11400714785074694791 + 9650029242287828579
	}

	h += n

	var i, sz = 0, len(b)
	for ; i+8 <= sz; i += 8 {
		k1 := round(0, u64(b[i:i+8:len(b)]))
		h ^= k1
		h = (h<<27|h>>37)*11400714785074694791 + 9650029242287828579
	}
	if i+4 <= sz {
		h ^= uint64(u32(b[i:i+4:len(b)])) * 11400714785074694791
		h = (h<<23|h>>41)*14029467366897019727 + 1609587929392839161
		i += 4
	}
	for ; i < sz; i++ {
		h ^= uint64(b[i]) * 2870177450012600261
		h = (h<<11 | h>>53) * 11400714785074694791
	}

	h ^= h >> 33
	h *= 14029467366897019727
	h ^= h >> 29
	h *= 1609587929392839161
	h ^= h >> 32

	return h
}

func round(a, v uint64) uint64 {
	a += v * 14029467366897019727
	return (a<<31 | a>>33) * 11400714785074694791
}

// from binary.LittleEndian.Uint64 package
func u64(b []byte) uint64 {
	_ = b[7] // bounds check hint to compiler; see golang.org/issue/14808
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
}

// from binary.LittleEndian.Uint32 package
func u32(b []byte) uint32 {
	_ = b[3] // bounds check hint to compiler; see golang.org/issue/14808
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}
