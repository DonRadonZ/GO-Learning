package main

import (
	"errors"
	"fmt"
	"os"
)

// Goals
// 1) Validate user input
// 	=> Show error message & exit if invalid input is provided
// 	- No negative numbers
// 	- Not 0
// 2) Store calculated results into file

const resultsFile = "result.txt"

func writeResultToFile(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.1f\n",ebt, profit, ratio)
	os.WriteFile(resultsFile, []byte(results), 0644)

}

func main() {

	revenue, err := getUserInput("your revenue is ")

	if err != nil {
		
		fmt.Println(err)
		
		// panic("Can't continue, sorry.")
		return
	}

	expense, err := getUserInput("your expense is ")

	if err != nil {
		
		fmt.Println(err)
		
		// panic("Can't continue, sorry.")
		return
	}


	tax_rate,err := getUserInput("Tax Rate is ")

	if err != nil {
		
		fmt.Println(err)
		
		// panic("Can't continue, sorry.")
		return
	}

	// if err1 != nil || err2 != nil || err3 != nil {
		
	// 	fmt.Println(err1)
		
	// 	// panic("Can't continue, sorry.")
	// 	return
	// }


	EBTValue, ProfitValue, RatioValue := calculatorValue(revenue, expense, tax_rate)

	formattedEBT := fmt.Sprintf("Earned before Tax Value: %.2f\n", EBTValue)
	formattedProfit := fmt.Sprintf("Earned after Tax Value: %.2f\n", ProfitValue)
	formattedRatio := fmt.Sprintf("Ratio: %.2f\n", RatioValue)

	fmt.Print(formattedEBT, formattedProfit, formattedRatio)
	writeResultToFile(EBTValue, ProfitValue, RatioValue)
}

func getUserInput(text string) (float64, error) {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("invalid amount. must be greater than 0")

	}

	return userInput, nil
}

func calculatorValue(revenue, expense, tax_rate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expense
	profit = float64(ebt) * (1 - tax_rate/100)
	ratio = ebt / profit
	return ebt, profit, ratio
}
