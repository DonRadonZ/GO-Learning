package main

import (
	"fmt"
)

func main() {


	revenue := getUserInput("your revenue is ")

	expense := getUserInput("your expense is ")

	tax_rate := getUserInput("Tax Rate is ")

	EBTValue, ProfitValue, RatioValue := calculatorValue(revenue, expense, tax_rate)

	formattedEBT := fmt.Sprintf("Earned before Tax Value: %.2f\n", EBTValue)
	formattedProfit := fmt.Sprintf("Earned after Tax Value: %.2f\n", ProfitValue)
	formattedRatio := fmt.Sprintf("Ratio: %.2f\n", RatioValue)

	fmt.Print(formattedEBT, formattedProfit, formattedRatio)

}

func getUserInput(text string) float64 {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)
	return userInput
}

func calculatorValue(revenue, expense, tax_rate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expense
	profit = float64(ebt) * (1 - tax_rate/100)
	ratio = ebt / profit
	return ebt, profit, ratio
}
