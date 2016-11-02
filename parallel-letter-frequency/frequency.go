package letter

// FreqMap is a map of runes to int
type FreqMap map[rune]int

// Frequency tallies the occurences of letters
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency runs Frequency concurrently
func ConcurrentFrequency(docs []string) FreqMap {
	ch := make(chan FreqMap)
	tot := FreqMap{}
	for _, doc := range docs {
		go func(d string) {
			ch <- Frequency(d)
		}(doc)
	}
	for range docs {
		for k, v := range <-ch {
			tot[k] += v
		}
	}
	return tot
}
