package pkg

import "math"

func MaxProfit(prices []int) int {
	min_price, max_price := math.MaxInt32, 0
	for p := range prices {
		if p < min_price {
			min_price = p
		} else if p-min_price > max_price {
			max_price = p - min_price
		}
	}
	return max_price
}
