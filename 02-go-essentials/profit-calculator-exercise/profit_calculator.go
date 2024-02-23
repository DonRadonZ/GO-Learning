package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expense float64
	var tax_rate float64

	fmt.Print("your revenue is ")
	fmt.Scan(&revenue)

	fmt.Print("your expense is ")
	fmt.Scan(&expense)

	fmt.Print("Tax Rate is ")
	fmt.Scan(&tax_rate)

	ebt := revenue - expense
	profit := float64(ebt) * (1 - tax_rate/100)
	ratio := ebt / profit

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)

}
