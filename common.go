package puzzle

// Solution is a function that can be called to produce a Euler puzzle solution
type Solution func() int

// IsPythagorean is intended to tell you whether these numbers make a
// pythagorean triplet
func IsPythagorean(a, b, c int) bool {
	return false
}

// IsPrime is intended to tell you whether a number is a prime
func IsPrime(n int) bool {
	return false
}

// IsPalindromic is intended to tell you whether to number is reversible using
// a base 10 representation
func IsPalindromic(n int) bool {
	return false
}

// Max is like math.Max, but works with int values
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
