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


func playGame() {
	horses :=  []Horse{
		{name: "Norsey", age: 4},
		{name: "Worsey", age: 1},
		{name: "Korsey", age: 2},
		{name: "Forsey", age: 6},
		{name: "Rob", age: 2},
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

}

func main() {
	balance := 100

	playGame()

	

	for balance > 0 {
		var input string
		fmt.Printf("Your balance is: $%d ", balance) 
		fmt.Println("Enter your bet (q to quit): ") 
		
		fmt.Scanln(&input)

		if input == "q"{
			fmt.Println("Bye!!")
			break
		}

		bet, err := strconv.Atoi(input) 
		if err != nil {
			fmt.Println("Invaild Input, Try again.")
			continue
		}

		if balance > bet {
			fmt.Println("Which one of Horsey's Cousins would you like to bet on?")

			var chosenHorse string
			fmt.Scanln(chosenHorse)
			
		}
	}
}

