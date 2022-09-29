package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	bytes := []byte(s)
	bytesMap := make(map[byte]byte)
	i := 0
	j := 0
	for j < len(bytes) {
		for i <= j {
			value, isExist := bytesMap[bytes[j]]
			if isExist && value == 1 {
				bytesMap[bytes[i]] = 0
				i++
				continue
			}
			bytesMap[bytes[j]] = 1
			if j-i+1 > maxLen {
				maxLen = j - i + 1
			}
			break
		}
		j++
	}

	return maxLen
}

func main() {
	s := "pwwkew"

	fmt.Println(lengthOfLongestSubstring(s))
}
