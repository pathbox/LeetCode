package LeetCode151

import "strings"

func reverseWords(s string) string {
	list := strings.Split(s, " ")
	var res []string
	for i := len(list) - 1; i >= 0; i-- {
		if len(list[i]) > 0 {
			res = append(res, list[i])
		}
	}
	s = strings.Join(res, " ")
	return s
}
