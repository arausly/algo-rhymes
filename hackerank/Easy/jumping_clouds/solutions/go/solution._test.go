package jumpingclouds

import (
	"fmt"
	"testing"
)

func TestJumpingOnClouds(t *testing.T) {
	ans := []int32{4, 3, 6, 28, 53, 54, 1, 46, 1}
	params := [][]int32{
		[]int32{0, 0, 1, 0, 0, 1, 0},
		[]int32{0, 0, 0, 1, 0, 0},
		[]int32{0, 0, 1, 0, 0, 0, 0, 1, 0, 0},
		[]int32{
			0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0,
			0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 1,
			0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 0, 0, 1, 0, 1, 0, 0,
		},
		[]int32{
			0, 1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0,
			0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1,
			0, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0,
			1, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1,
			0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 1,
			0, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 1,
			0, 0, 0, 0, 0, 0, 1, 0, 0, 0,
		},
		[]int32{
			0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0,
			0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0,
			0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1,
			0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 1,
			0, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0,
			0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1,
			0, 1, 0, 1, 0, 0, 0, 0, 1, 0,
		},
		[]int32{0, 0},
		[]int32{
			0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1,
			0, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0,
			1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0,
			0, 1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 0, 1,
			0, 1, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 1,
			0, 0, 0, 1, 0, 0, 1, 0, 1, 0,
		},
		[]int32{0, 1, 0},
	}

	fmt.Printf("------------------------Jumping Clouds-----------------------\n")

	for i := 0; i < len(ans); i++ {
		j := jumpingOnClouds(params[i])
		if j != ans[i] {
			t.Errorf("Expected %v but received %v", ans[i], j)
		}
	}
}
