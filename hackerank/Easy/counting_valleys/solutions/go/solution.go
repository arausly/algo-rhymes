package countingvalleys

import (
	"strings"
)

func countingValleys(n int, paths string) int {
	valleyCount, depthTracker, seaLevel := 0, 0, 0
	s := strings.Split(paths, "")
	for _, path := range s {
		if path == "U" {
			seaLevel++
		} else {
			seaLevel--
		}

		if seaLevel == -1 {
			depthTracker = seaLevel
		}

		if seaLevel == 0 && depthTracker == -1 {
			valleyCount++
			depthTracker = 0
		}
	}
	return valleyCount
}
