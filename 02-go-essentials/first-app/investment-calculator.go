package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {
	var investmentAmount float64
var years float64
var expectedReturnRate = 5.5
	

	// fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	//fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for Information)\n: %.2f", futureRealValue)
	// output information
	// fmt.Println("Future Value: ",futureValue)
	// fmt.Printf(`Future Value: %.2f\n

	// Future Value (adjusted for Information): %.2f`, futureValue, futureRealValue)
	// fmt.Println("Future Value (adjusted for Inflation): ",futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string){
	fmt.Print(text)
}


func calculateFutureValues(){
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years) 
   rfv := futureValue / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}