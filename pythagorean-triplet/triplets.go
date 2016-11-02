// Package pythagorean does some magic
package pythagorean

// HashMap is a map of ints to ints
type HashMap map[int]int

// Triplet is a slice of ints
type Triplet []int

// Range given min and max, finds triplets within min and max
func Range(min, max int) []Triplet {
	var allTriplets []Triplet
	hashmap := make(HashMap, (max - min + 1))
	// create hash maps to save calculating powers each time
	for i := min; i <= max; i++ {
		hashmap[i] = i * i
	}
	// slow but can add concurrency
	for p1 := min; p1 <= max-2; p1++ {
		for p2 := p1 + 1; p2 <= max-1; p2++ {
			for p3, v := range hashmap {
				if hashmap[p1]+hashmap[p2] == v {
					allTriplets = append(allTriplets, Triplet{p1, p2, p3})
				}
			}
		}
	}
	return allTriplets
}

// Sum finds given num, triplets less than num that sum to num
func Sum(sum int) []Triplet {
	var allTriplets []Triplet
	triplets := Range(1, sum)
	for _, v := range triplets {
		if v[0]+v[1]+v[2] == sum {
			allTriplets = append(allTriplets, v)
		}
	}
	return allTriplets
}
