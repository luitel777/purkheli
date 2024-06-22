package utils

// removes trailing slash
func RemoveTrail(s string) string {
	if len(s) == 1 {
		return s
	}
	if string(s[len(s)-1]) == "/" {
		return s[:len(s)-1]
	}
	return s
}
