package series

func All(n int, s string) []string {
	res := make([]string, 0)
	for i := 0; i <= len(s)-n; i++ {
		res = append(res, s[i:i+n])
	}
	return res
}

func UnsafeFirst(n int, s string) string {
	if n > len(s) {
		return s
	}
	return s[0:n]
}
