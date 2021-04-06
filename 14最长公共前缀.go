package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	stop := true
	start := 1
	prefix := ""
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	for stop {
		for k, v := range strs {
			if start > len(v) {
				stop = false
				if k > 0 {
					prefix = prefix[0 : len(prefix)-1]
				}
				break
			}
			if k == 0 {
				prefix = v[0:start]
			}
			if v[0:start] != prefix {
				stop = false
				if start == 1 {
					prefix = ""
				} else {
					prefix = prefix[0 : len(prefix)-1]
				}
				break
			}
		}
		start += 1
	}
	return prefix
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}
