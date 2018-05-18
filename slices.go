package utils

import ()

// returns index where predicate is true
// USAGE:
//   Find(len(slice), func(i int) bool { return slice[i] == element })
func Find(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}

	return -1
}
