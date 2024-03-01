package main

import (
	"fmt"
	"examples.com/bank/bank/fileops"
)

	const accountBalanceFile = "balance.txt"



func main() {
	var accountBalance,err = fileops.GetBalanceFromFile(accountBalanceFile)

		

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------")
		// panic("Can't continue, sorry.")
	}

	fmt.Println("Welcome to Go Bank!")

	for {
		PresentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		// wantCheckBalance := choice == 1

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
			fileops.GetBalanceFromFile(accountBalanceFile)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalia amount. Must be greater than 0.")
				// return
				continue
			}

			accountBalance += depositAmount // accountBalance = accountBalance + depositAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
			fileops.WriteBalanceToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Print("Your withdraw: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				return
			}
			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				return
			}
			accountBalance -= withdrawalAmount // accountBalance = accountBalance + withdrawalAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
			fileops.WriteBalanceToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for using our service.")
			return
			//break
		}
	}
}



