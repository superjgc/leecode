package _413_拥有最多糖果的孩子

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0
	for i := 0; i < len(candies); i++ {
		max = func(left int, right int) int {
			if left > right {
				return left
			}
			return right
		}(max, candies[i])
	}
	var result = make([]bool, len(candies))
	for i := range candies {
		if candies[i]+extraCandies >= max {
			result[i] = true
		}
	}
	return result
}
