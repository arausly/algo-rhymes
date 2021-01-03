package repeatedstrings

import (
	"math"
	"strings"
)

func countA(s string, f *int64) {
	for _, val := range s {
		if val == 'a' {
			*f++
		}
	}
}

//accepted and most optimal solution ✅✅
func repeatedString(s string, n int64) int64 {
	var f int64
	l := int64(len(s))

	switch {
	case !strings.Contains(s, "a"):
		return 0
	case l == 1:
		return n
	case l < n:
		g := math.Floor(float64(n / l))
		m := n % l
		countA(s, &f)
		f *= int64(g)

		if m > 0 {
			countA(s[:m], &f)
		}
		return f
	default:
		countA(s[:n], &f)
		return f
	}
}

//⚠⚠⚠⚠⚠⚠⚠⚠⚠⚠ WARNING!! ⚠⚠⚠⚠⚠⚠⚠⚠⚠⚠⚠⚠⚠⚠
//valid solution but NOT optimal solution.
//doesn't pass tests with time-limits on them
func repeatedString2(s string, n int64) int64 {
	var f int64
	r, l := s, int64(len(s))

	if !strings.Contains(s, "a") {
		return 0
	}

	if l == 1 {
		return n
	}

	//build the repeated string
	if l < n {
		for int64(len(r)) < n {
			d := n - int64(len(r))
			if int(d) > len(s) {
				d = int64(len(s))
			}
			r += s[:d]
		}
	}

	//count frquencies of a in repeated string
	for _, val := range r[:n] {
		if val == 'a' {
			f++
		}
	}

	return f
}
