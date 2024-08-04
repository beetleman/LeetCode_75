package leetcode75

func isSubsequence(s string, t string) bool {
	checked := 0
	if s == "" {
		return true
	}
	for i := 0; i < len(t); i++ {
		if t[i] == s[checked] {
			checked++
			if checked == len(s) {
				return true
			}
		}
	}
	return false
}
