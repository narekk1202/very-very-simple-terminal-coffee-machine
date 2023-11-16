package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

type coffes struct {
	name  string
	price int
}

var coffeeOptions = []coffes{
	{"Espresso", 2},
	{"Cappuccino", 3},
	{"Latte", 4},
	{"Mocha", 5},
	{"Americano", 2},
}

func main() {
	for {
		var selectedCoffee int
		var sum int
		var coffee coffes
		fmt.Println("Choose your coffee!")
		fmt.Println("========================")
		printCoffees()

		fmt.Print("Enter your choice: ")
		fmt.Scan(&selectedCoffee)
		clearTerminal()
		if selectedCoffee > 0 && selectedCoffee <= len(coffeeOptions) {
			coffee = coffeeOptions[selectedCoffee-1]
		}

		if len(coffee.name) > 0 {
			fmt.Print("Enter the sum: ")
			fmt.Scan(&sum)
			if sum == coffee.price {
				fmt.Println("Your order is", coffee.name)
				askForSugar()
			} else if sum > coffee.price {
				fmt.Printf("Here is your change $%d\n", sum-coffee.price)
				fmt.Println("Your order is", coffee.name)
				askForSugar()
			} else {
				fmt.Println("Not enough!")
			}
		} else {
			fmt.Println("Invalid choice!")
			time.Sleep(1 * time.Second)
			clearTerminal()
		}
	}
}

func printCoffees() {
	for key, coffee := range coffeeOptions {
		fmt.Printf("%d. %s $%d\n", key+1, coffee.name, coffee.price)
	}
}

func askForSugar() {
	var acceptSugar int
	var spoons int

	fmt.Print("Do you want sugar? 1. Yes 2. No: ")
	fmt.Scan(&acceptSugar)
	clearTerminal()

	if acceptSugar == 1 {
		fmt.Print("How many spoons of sugar you want: ")
		fmt.Scan(&spoons)
		clearTerminal()
	}

	if spoons != 0 && spoons <= 5 {
		fmt.Println("Wait for your coffee!")
		time.Sleep(time.Duration(spoons) * time.Second)
		fmt.Println("Your coffee is ready!")
		time.Sleep(3 * time.Second)
		clearTerminal()
	} else if acceptSugar == 2 {
		fmt.Println("Here is your coffee, enjoy!")
		time.Sleep(3 * time.Second)
		clearTerminal()
	}
}

func clearTerminal() {
	screen.Clear()
	screen.MoveTopLeft()
}
