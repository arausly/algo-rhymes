package repeatedstrings

import "testing"

func TestRepeatedStrings(t *testing.T) {
	var (
		n1 int64
		n2 int64
	)

	s1, s2, n1, n2 := "aba", "a", 10, 1000000000000

	t1, t2 := repeatedString(s1, n1), repeatedString(s2, n2)

	if t1 != 7 {
		t.Errorf("Expected 7 but received %v", t1)
	}

	if t2 != 1000000000000 {
		t.Errorf("Expected 1000000000000 but received %v", t2)
	}
}
