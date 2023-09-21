package leet

func MaxProfit(prices []int) int {
	buy := 0
	var sell int
	sell = buy
	l := len(prices)
	max_profit := 0
	var prev_max int
	for i := 0; i < l; i++ {

		if prices[i] < prices[buy] {
			buy = i
			sell = buy

		}
		if prices[i] > prices[sell] {
			sell = i
			prev_max = prices[sell] - prices[buy]
			if prev_max > max_profit {
				max_profit = prev_max
			}

		}
	}
	return max_profit
}
