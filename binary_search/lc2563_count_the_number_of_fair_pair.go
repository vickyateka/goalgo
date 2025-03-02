

func countFairPairs(arr []int, lower int, upper int) int64 {
    n := len(arr)
    sort.Slice(arr, func(i, j int) bool {
        return arr[i] < arr[j]
    })

    lower_bound := func(v []int, l, r, val int) int{
        res := -1
        for l <= r{
            m := (l + r) / 2
            if v[m] >= val {
                res = m
                r = m - 1
            }else{
                l = m + 1
            }
        }

        return res
    }

    upper_bound := func(v []int, l, r, val int) int{
        res := -1
        for l <= r{
            m := (l + r) / 2
            if v[m] <= val {
                res = m
                l = m + 1
            }else{
                r = m - 1
            }
        }

        return res
    }

    var ans int64
    for i := 0; i < n; i ++{
        y := upper_bound(arr, i + 1, n - 1, upper - arr[i])
        x := lower_bound(arr, i + 1, n - 1, lower - arr[i])

        if x >= 0 && y >= 0 {
            ans += int64(y - x + 1)
        }
    }

    return ans
}

// 0, 1, 2, 3, 4, 5
// 0, 1, 4, 4, 5, 7
// l = 3, u = 6

// b = lower_bound(u - a[i])
// a = lower_bound(l + a[i])
// ans += (b - a + 1)

func main(){

    // call the function with the actual input values. 
    // take over from here
}
