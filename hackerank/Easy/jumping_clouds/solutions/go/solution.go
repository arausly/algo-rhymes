package jumpingclouds

//Initial thought process
func jumpingOnClouds1(c []int32) int32 {
	var v []int32
	//get valid indices
	for index, val := range c {
		if val == 0 {
			v = append(v, int32(index))
		}
	}

	counter, m := 0, []int32{v[0]}
	//get minimum jumps
	for counter+2 <= len(v) {
		n1, n2 := v[counter], v[counter+1]

		if n2-n1 == 1 && counter+2 <= len(v)-1 {
			n3 := v[counter+2]
			if n3-n1 == 2 {
				m = append(m, n3)
				counter += 2
			} else {
				m = append(m, n2)
				counter++
			}
		} else {
			m = append(m, n2)
			counter++
		}
	}
	return int32(len(m)) - 1
}

//second thought process
func jumpingOnClouds2(c []int32) int32 {
	counter, m := 0, []int32{0}

	//get minimum jumps
	for counter+2 <= len(c) {
		n1, n2 := counter, counter+1
		if c[counter] == 0 {
			if n2-n1 == 1 && counter+2 <= len(c)-1 && c[counter+2] == 0 {
				n3 := counter + 2
				if n3-n1 == 2 {
					m = append(m, int32(n3))
					counter += 2
				} else {
					m = append(m, int32(n2))
					counter++
				}
			} else {
				m = append(m, int32(n2))
				counter++
			}
		}
	}
	return int32(len(m)) - 1
}

//accepted and most optimal solution ✅✅
func jumpingOnClouds(c []int32) (j int32) {
	i := 0
	for {
		if i+2 < len(c) {
			if c[i+2] == 0 {
				j++
				i += 2
			} else if c[i+1] == 0 {
				j++
				i++
			}
		} else if i+1 < len(c) {
			j++
			i++
		} else {
			break
		}
	}
	return j
}
