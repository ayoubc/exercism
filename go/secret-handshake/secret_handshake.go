package secret

func Handshake(n uint) []string {
	shakes := map[uint]string{
		1: "wink",
		2: "double blink",
		4: "close your eyes",
		8: "jump",
	}
	res := make([]string, 0)
	for i := 0; i < 4; i++ {
		c := n & (1 << i)
		if c != 0 {
			res = append(res, shakes[c])
		}
	}

	// reverse
	if n&(1<<4) != 0 {
		reverse(res)
	}
	return res
}

func reverse(a []string) {
	n := len(a)
	for i := 0; i < n/2; i++ {
		a[i], a[n-i-1] = a[n-i-1], a[i]
	}
}
