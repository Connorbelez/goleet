package working

import "math"

func checkOk(piles []int, hours int, rate int) bool {
	for _, val := range piles {
		hours -= int(math.Ceil(float64(val) / float64(rate)))
	}
	return hours >= 0
}
func minEatingSpeed(piles []int, h int) int {

	max1 := func(arr []int) int {

		max := arr[0]

		for _, val := range arr {

			if max < val {
				max = val
			}
		}
		return max
	}(piles)

	if h == len(piles) {
		return max1
	}
	if len(piles) == 1 {
		return int(math.Ceil(float64(piles[0]) / float64(h)))
	}

	smallest := max1

	lp := 1
	rp := max1

	for lp <= rp {
		med := (lp + rp) / 2

		if checkOk(piles, h, med) {
			if med < smallest {
				smallest = med
			}
			rp = med - 1
		} else {
			lp = med + 1
		}

	}
	return smallest
}
