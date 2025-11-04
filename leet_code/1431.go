package main

func max(array []int) int {
	max := array[0]
	for i := 1; i < len(array); i++ {
		if array[i] > max {
			max = array[i]
		}
	}
	return max
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := max(candies)
	result := make([]bool, len(candies))
	for i := 0; i < len(candies); i++ {
		if candies[i] + extraCandies >= max {
			result[i] = true
		} else {
			result[i] = false
		}
	}
	return result
}