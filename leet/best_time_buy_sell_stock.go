package leet

func MaxProfit(prices []int) int {
	buy := 0
	var sell int
	sell = buy
	l := len(prices)
	max_profit := 0
	for i := 0; i < l; i++ {

		if prices[i] < prices[buy] {
			buy = i
			sell = buy

		}
		if prices[i] > prices[sell] {
			sell = i
			if prices[sell]-prices[buy] > max_profit {
				max_profit = prices[sell] - prices[buy]
			}

		}
	}
	return max_profit
}
