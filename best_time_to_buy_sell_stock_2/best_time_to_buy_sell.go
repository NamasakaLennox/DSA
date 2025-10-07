func maxProfit(prices []int) int {
    buy, profit := prices[0], 0
    for _, val := range(prices) {
        if (val < buy) {
            buy = val
        } else {
            profit += val - buy
            buy = val
        }
    }
    return profit
}
