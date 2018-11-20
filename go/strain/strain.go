package strain

// Ints is a collection of ints
type Ints []int

// Lists is a collection int arrays
type Lists [][]int

// Strings is a collection of strings
type Strings []string

// Keep filters a collection of Ints
// where the predicate function returns true
func (col Ints) Keep(keep func(int) bool) Ints {
	var keeps []int
	for _, e := range col {
		if keep(e) {
			keeps = append(keeps, e)
		}
	}
	return keeps
}

// Discard discards every element in a collection of Ints
// where the predicate function returns true
func (col Ints) Discard(discard func(int) bool) Ints {
	keep := func(e int) bool {
		return !discard(e)
	}
	return col.Keep(keep)
}

// Keep keeps every element in a collection of Ints where
// the predicate function returns true
func (col Strings) Keep(keep func(string) bool) Strings {
	var keeps Strings
	for _, e := range col {
		if keep(e) {
			keeps = append(keeps, e)
		}
	}
	return keeps
}

// Discard elements when the predicate returns true
func (col Strings) Discard(discard func(string) bool) Strings {
	keep := func(e string) bool {
		return !discard(e)
	}
	return col.Keep(keep)
}

// Keep elements when the predicate returns true
func (col Lists) Keep(keep func([]int) bool) Lists {
	var keeps [][]int
	for _, e := range col {
		if keep(e) {
			keeps = append(keeps, e)
		}
	}
	return keeps
}
