package strings

func lengthOfLongestSubstring(s string) int {

	n := len(s)
	j, ans := 0, 0

	freq := make(map[byte]int)

	for i := 0; i < n; i++ {
		freq[s[i]]++
		for freq[s[i]] > 1 {
			freq[s[j]]--
			j++
		}
		if ans < (i - j + 1) {
			ans = i - j + 1
		}
	}

	return ans
}

//func main(){
//
//
//}
