package countingvalleys

import "testing"

func TestCountingValleys(t *testing.T) {
	path1, steps1, paths2, step2 := "DDUUDDUDUUUD", 12, "UDDDUDUU", 8
	valleys1, valleys2 := countingValleys(steps1, path1), countingValleys(step2, paths2)

	if valleys1 != 2 {
		t.Errorf("Expected 2 but received %v", valleys1)
	}

	if valleys2 != 1 {
		t.Errorf("Expected 1 but received %v", valleys2)
	}
}
