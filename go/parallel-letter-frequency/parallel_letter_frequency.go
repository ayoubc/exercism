package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// var wg sync.WaitGroup

func (f *FreqMap) update(other FreqMap) {
	for key, value := range other {
		(*f)[key] += value
	}
}

func ConcurrentFrequency(phrases []string) FreqMap {
	res := FreqMap{}
	stream := make(chan FreqMap, len(phrases))

	go func() {
		defer close(stream)
		for _, phrase := range phrases {
			stream <- Frequency(phrase)
		}
	}()

	for f := range stream {
		res.update(f)
	}

	return res
}
