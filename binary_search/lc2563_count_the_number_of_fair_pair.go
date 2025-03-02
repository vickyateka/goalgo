package main

import (
	"sort"
)

func lowerBound(v []int, l, r, val int) int {
	for l <= r {
		m := (l + r) / 2
		if v[m] >= val {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}

func upperBound(v []int, l, r, val int) int {
	for l <= r {
		m := (l + r) / 2
		if v[m] > val {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}

// Count fair pairs
func countFairPairs(arr []int, lower int, upper int) int64 {
	n := len(arr)
	sort.Ints(arr) // Sort the array

	var ans int64
	for i := 0; i < n; i++ {
		x := lowerBound(arr, i+1, n-1, lower-arr[i])
		y := upperBound(arr, i+1, n-1, upper-arr[i]) - 1

		if x <= y {
			ans += int64(y - x + 1)
		}
	}
	return ans
}

//
//func main() {
//	// Example test case
//	arr := []int{1, 2, 3, 4, 5}
//	lower := 2
//	upper := 5
//
//	result := countFairPairs(arr, lower, upper)
//	fmt.Println("Count of Fair Pairs:", result)
//}
