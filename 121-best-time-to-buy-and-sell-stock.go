// You are given an array prices where prices[i] is the price of a given stock on the ith day.
//
// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
//
// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
//
//
//
// Example 1:
//
// Input: prices = [7,1,5,3,6,4]
// Output: 5
// Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
// Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
// Example 2:
//
// Input: prices = [7,6,4,3,1]
// Output: 0
// Explanation: In this case, no transactions are done and the max profit = 0.

package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	bestSell := 0
	bestBuy := prices[0]

	for _, price := range prices {
		if price < bestBuy {
			bestBuy = price
		}
		bestValue := price - bestBuy
		if bestValue > bestSell {
			bestSell = bestValue
		}
	}

	return bestSell
}

func main() {
	prices1 := []int{7, 1, 5, 3, 6, 4}
	prices2 := []int{7, 6, 4, 3, 1}
	prices3 := []int{2, 1, 2, 0, 1}
	result1 := maxProfit(prices1)
	result2 := maxProfit(prices2)
	result3 := maxProfit(prices3)
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}
