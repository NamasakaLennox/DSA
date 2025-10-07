func maxProfit(prices []int) int {
    lowest, profit := prices[0], 0
    for _, val := range(prices) {
        if (val < lowest) {
            lowest = val
        } else if (val - lowest > profit) {
            profit = val - lowest
        }
    }
    return profit
}
