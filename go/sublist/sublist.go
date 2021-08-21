package sublist

func Sublist(a, b []int) string {
	c1, c2 := isSublist(a, b), isSublist(b, a)
	if c1 && c2 {
		return "equal"
	} else if c1 {
		return "sublist"
	} else if c2 {
		return "superlist"
	} else {
		return "unequal"
	}
}

func isSublist(a, b []int) bool {
	if len(b) < len(a) {
		return false
	}
	for i := 0; i <= len(b)-len(a); i++ {
		ok := true
		for j := i; j < i+len(a); j++ {
			if a[j-i] != b[j] {
				ok = false
				break
			}
		}
		if ok {
			return true
		}
	}
	return false
}
