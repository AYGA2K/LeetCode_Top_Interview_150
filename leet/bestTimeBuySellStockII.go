package leet

func MaxProfitII(prices []int) int {
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
			max_profit = max_profit + prices[sell] - prices[buy]
			buy = i
		}
	}
	return max_profit
}
