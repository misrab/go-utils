package utils

import ()

func FlipMapUint64(original map[uint64]uint64) map[uint64]uint64 {
	flipped := make(map[uint64]uint64)

	for k, v := range original {
		flipped[v] = k
	}

	return flipped
}
