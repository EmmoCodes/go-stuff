package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile, 1000)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------")
		// panic("Can't continue, sorry!")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())

	for {
		presentOptions()

		var choice int
		print("Your Choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)

		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. must be greater than 0.")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)

		case 3:
			fmt.Print("Your Withdrawal amount: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. must be greater than 0.")
				return
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount. You cant withdraw more than you have.")
				return
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)

		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank!")
			return
		}
	}

}
