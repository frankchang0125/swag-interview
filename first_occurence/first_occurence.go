package first_occurence

import (
	"strings"
)

func firstOccurence(s string, x string) int32 {
	targets := strings.Split(x, "*")

	if len(targets) == 1 {
		// Without "*"
		for i := 0; i <= len(s) - len(x); i++ {
			if s[i:i+len(x)] == x {
				return int32(i)
			}
		}
	} else if len(targets) == 2 {
		// With "*"
		i := 0

		for i <= len(s) - len(targets[0]) {
			for ; i <= len(s) - len(targets[0]); i++ {
				if s[i:i+len(targets[0])] == targets[0] {
					break
				}
			}

			j := i + len(targets[0]) + 1
			if j <= len(s) - len(targets[1]) {
				if s[j:j+len(targets[1])] == targets[1] {
					return int32(i)
				}
			}

			i++
		}
	} 

	return -1
}
