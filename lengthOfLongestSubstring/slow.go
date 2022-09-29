package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	bytes := []byte(s)

	for curLen := 1; curLen < 128; curLen++ {
		for i := 0; i+curLen-1 < len(bytes); i++ {
			isOK := true
			j := i + curLen - 1

			bytesMap := make(map[byte]byte)
			for k := i; k <= j; k++ {
				value, isExist := bytesMap[bytes[k]]
				if isExist && value == 1 {
					isOK = false
					break
				}
				bytesMap[bytes[k]] = 1
			}

			if isOK {
				maxLen = curLen
				break
			}
		}

		if maxLen != curLen {
			break
		}
	}

	return maxLen
}

func main() {
	s := " "

	fmt.Println(lengthOfLongestSubstring(s))
}
