package strain

// Collections, hm?  For this exercise in Go you'll work with slices as
// collections.  Define the following in your solution:
//
//    type Ints []int
//    type Lists [][]int
//    type Strings []string
//
// Then complete the exercise by implementing these methods:
//
//    (Ints) Keep(func(int) bool) Ints
//    (Ints) Discard(func(int) bool) Ints
//    (Lists) Keep(func([]int) bool) Lists
//    (Strings) Keep(func(string) bool) Strings

type Ints []int
type Lists [][]int
type Strings []string

// Keep lkjklj
func Keep(input []int) []int {
	var even []int

	for _, num := range input {
		if num%2 == 0 {
			even = append(even, num)
		}
	}
	return even
}
