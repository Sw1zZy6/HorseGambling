package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"

)

type Horse struct {
	name string
	age int 
	time float64
}

func roundTime(value float64) float64 {
	return float64(int(value*1000+0.5)) / 1000
}

func playGame(balance float64, bet float64) float64{
	horses :=  []Horse{
		{name: "norsey", age: 4},
		{name: "worsey", age: 1},
		{name: "korsey", age: 2},
		{name: "forsey", age: 6},
		{name: "rob", age: 2},
	}

	list := [5]string{
		"5th", 
		"4th",
		"3rd",
		"2nd",
		"1st",
	}

	for i := range horses {
		horses[i].time = roundTime(rand.Float64())
	}

	sort.Slice(horses, func(i int, j int) bool {
		return horses[i].time < horses[j].time
	})

	for i, horse := range horses {
		fmt.Println(list[i], ". ", horse.name, horse.time)
	}

	fmt.Println("\nWhich one of Horsey's Cousins would you like to bet on?")
	for i := 0; i < len(horses); i++ {
		fmt.Println(horses[i].name)
	}

	var chosenHorse string
	fmt.Scanln(&chosenHorse)
	
	checkingChosenHorse := 0

	for i := 0; i < len(horses); i++ {
		if chosenHorse == horses[i].name{
		} else {
			checkingChosenHorse++
		}
	}

	if checkingChosenHorse == 5 {
		fmt.Println("Thats not an option, try again.")
	} else {
		fmt.Println("You have chosen: ", chosenHorse)
	
		if (chosenHorse == horses[0].name) {
			balance += bet * 2
		} else if (chosenHorse == horses[1].name) {
			balance += bet 
		} else if (chosenHorse == horses[3].name) {
			balance -= bet 
		} else if (chosenHorse == horses[4].name) {
			balance -= bet * 2
		}
	}
	return balance
}

func main() {
	balance := 100.0

	for balance > 0 {
		var input string
		fmt.Println("")
		fmt.Printf("Your balance is: $%.2f ", balance) 
		fmt.Println("Enter your bet (q to quit): ") 
		
		fmt.Scanln(&input)

		if input == "q"{
			fmt.Println("Bye!!")
			break
		}

		bet, err := strconv.ParseFloat(input, 64)
		
		if err != nil {
			fmt.Println("Invaild Input, Try again.")
			continue
		}

		if bet > balance {
			fmt.Println("Insufficent Funds.")
		} else {}
	
		balance = playGame(balance, bet)
	}
}

