package lists

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}

	fmt.Println(prices[0:1])
	prices[1] = 9.99

	prices = append(prices, 5.99)
	prices = prices[1:]
	fmt.Println(prices)

	discountPrices := []float64{10.99, 80.99, 20.59}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}

// func main() {
// 	var productNames []string =  []string{"Book","Beef","Tea","Mackerel"}
// 	prices := []float64{25, 40, 2.5, 99.25}
// 	fmt.Println(prices)

//     productNames[2] = "A Carpet"

// 	fmt.Println(productNames)

// 	fmt.Println(prices[3],productNames[1])

// 	featuredPrices := prices[1:]
// 	highlightedPrice := featuredPrices[:1]
// 	fmt.Println(highlightedPrice)
// }
