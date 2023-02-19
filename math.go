package common_util

import (
	"sort"
)

func Max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// SortSlice 小 -> 大 , [reverse] = true 大 -> 小
func SortSlice[S Ordered](s []S, reverse ...bool) {
	if len(reverse) > 0 && reverse[0] {
		sort.Slice(s, func(i, j int) bool {
			return s[i] > s[j]
		})
	} else {
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
	}
}
