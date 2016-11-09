package strain

// Ints ...
type Ints []int

// Lists ...
type Lists [][]int

// Strings ...
type Strings []string

// Keep ...
func (in Ints) Keep(f func(int) bool) Ints {
	if in == nil {
		return nil
	}
	res := Ints{}
	for _, v := range in {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}

// Discard ...
func (in Ints) Discard(f func(int) bool) Ints {
	if in == nil {
		return nil
	}
	res := Ints{}
	for _, v := range in {
		if !f(v) {
			res = append(res, v)
		}
	}
	return res
}

// Keep ...
func (in Lists) Keep(f func([]int) bool) Lists {
	res := Lists{}
	for _, v := range in {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}

// Keep ...
func (in Strings) Keep(f func(string) bool) Strings {
	res := Strings{}
	for _, v := range in {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}
