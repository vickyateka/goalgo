package arrays

// import "slices"
func productExceptSelf(arr []int) []int {
	n := len(arr)
	pref := make([]int, n)
	suf := make([]int, n)
	for i := 0; i < n; i++ {
		pref[i], suf[i] = 1, 1
	}
	for i := 1; i < n; i++ {
		pref[i] = arr[i-1] * pref[i-1]
	}
	for i := n - 2; i >= 0; i-- {
		suf[i] = suf[i+1] * arr[i+1]
	}
	var res []int
	for i := 0; i < n; i++ {
		if i == 0 {
			res = append(res, suf[i])
		} else if i == n-1 {
			res = append(res, pref[i])
		} else {
			res = append(res, suf[i]*pref[i])
		}
	}

	return res
}

//
//func main() {
//	arr := []int{1, 2, 3, 4}
//	result := productExceptSelf(arr)
//	fmt.Println(result)
//
//}
