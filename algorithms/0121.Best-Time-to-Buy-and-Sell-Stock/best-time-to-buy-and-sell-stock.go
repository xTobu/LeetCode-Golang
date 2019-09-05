package problem0121

import (
	"math"
)

// Input: [7,1,5,3,6,4]
// Output: 5
// Explanation:
// Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
// Not 7-1 = 6, as selling price needs to be larger than buying price.
// 獲利算法是 賣出 - 買入，故只能 後減前

// 以下使用單次遍歷
func maxProfit(prices []int) int {

	// minPrice, maxProfit
	// 用來儲存每次遍歷後找到的 最小價錢 minPrice, 最大獲利 maxProfit
	minPrice, maxProfit := math.MaxInt64, 0

	for _, price := range prices {
		// 使用 min 來找出當次和過去的最小價錢
		minPrice = min(minPrice, price)

		// 使用 max 來比較當次和過去的最大獲利
		// ps:
		// 因為 minPrice 只會來自過去報價(買入價錢)，
		// 故可直接用 price - minPrice 來取得當次的最大獲利
		maxProfit = max(maxProfit, price-minPrice)
	}

	return maxProfit
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
