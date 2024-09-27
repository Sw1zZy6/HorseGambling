package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

type Horse struct {
	name string
	age int 
	time float64
	experience float64
}

func roundTime(value float64) float64 {
	return float64(int(value*1000+0.5)) / 1000
}

func playGame(balance float64, bet float64) float64{
	horses :=  []Horse{
		{name: "norsey", age: 4, experience: 2,}, 
		{name: "worsey", age: 1, experience: 0,},
		{name: "korsey", age: 2, experience: 1,},
		{name: "forsey", age: 6, experience: 4},
		{name: "rob", age: 3, experience: 1},
	}

	list := [5]string{
		"5th", 
		"4th",
		"3rd",
		"2nd",
		"1st",
	}

	for i := range horses {
		horses[i].time = roundTime(rand.Float64()) + roundTime(horses[i].experience / 25)
	}

	sort.Slice(horses, func(i int, j int) bool {
		return horses[i].time < horses[j].time
	})

	fmt.Println("\nWhich one of Horsey's Cousins would you like to bet on?\n")
	fmt.Println("*Horses with lower experience have lower chance of winning, but higher profits.\n")
	for i := 0; i < len(horses); i++ {
		fmt.Println(horses[i].name)
		fmt.Println("Experence: ", horses[i].experience, "\n")
	}
	fmt.Println("")
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
		fmt.Println("\nYou have chosen: ", chosenHorse, "\n")
	
		/*if (chosenHorse == horses[4].name) {
			if horses[4].experience == 0{
				percent := bet / 2.5
				balance += bet * 2 + percent
			} else if horses[4].experience > 0 && horses[4].experience < 3 {
				percent := bet / 5
				balance += bet * 2 + percent
			} else if horses[4].experience >= 3 {
				percent := bet / 10
				balance += bet * 2 + percent
			}

		} else if (chosenHorse == horses[3].name) {
			if horses[3].experience == 0{
				percent := bet / 2.5
				balance += bet + percent
			} else if horses[3].experience > 0 && horses[3].experience < 3 {
				percent := bet / 5
				balance += bet + percent
			} else if horses[3].experience >= 3 {
				percent := bet / 10
				balance += bet + percent
			}

		} else if (chosenHorse == horses[1].name) {
			if horses[1].experience == 0{
				percent := bet / 2.5
				balance -= bet + percent
			} else if horses[1].experience > 0 && horses[1].experience < 3 {
				percent := bet / 5
				balance -= bet + percent
			} else if horses[1].experience >= 3 {
				percent := bet / 10
				balance -= bet + percent
			}

		} else if (chosenHorse == horses[0].name) {
			if horses[0].experience == 0{
				percent := bet / 2.5
				balance -= bet * 2 + percent
			} else if horses[0].experience > 0 && horses[4].experience < 3 {
				percent := bet / 5
				balance -= bet * 2 + percent
			} else if horses[0].experience >= 3 {
				percent := bet / 10
				balance -= bet * 2 + percent
			}
		}
		*/
	}

	for i, horse := range horses {
		fmt.Println(list[i], ". ", horse.name, horse.time)
		time.Sleep(1 * time.Second)
	}

	switch chosenHorse {
		case horses[0].name:
			if (chosenHorse == horses[0].name) {
				if horses[0].experience == 0{
					percent := bet / 2.5
					balance -= bet * 2 + percent
					fmt.Printf("%s %s %.2f", "\nYou've Lost: ", "$", bet + bet + percent)

				} else if horses[0].experience > 0 && horses[4].experience < 3 {
					percent := bet / 5
					balance -= bet * 2 + percent
					fmt.Printf("%s %s %.2f", "\nYou've Lost: ", "$", bet + bet + percent)

				} else if horses[0].experience >= 3 {
					percent := bet / 10
					balance -= bet * 2 + percent
					fmt.Printf("%s %s %.2f", "\nYou've Lost: ", "$", bet + bet + percent)
				} 
			}
		case horses[1].name:
			if (chosenHorse == horses[1].name) {
				if horses[1].experience == 0{
					percent := bet / 2.5
					balance -= bet + percent
					fmt.Printf("%s %s %.2f", "\nYou've Lost: ", "$", bet + percent)
				} else if horses[1].experience > 0 && horses[1].experience < 3 {
					percent := bet / 5
					balance -= bet + percent
					fmt.Printf("%s %s %.2f", "\nYou've Lost: ", "$", bet + percent)
				} else if horses[1].experience >= 3 {
					percent := bet / 10
					balance -= bet + percent
					fmt.Printf("%s %s %.2f", "\nYou've Lost: ", "$", bet + percent)
				}
			}
		case horses[2].name:
			fmt.Println("\nYou have not gain nor lost any money.")
		case horses[3].name: 
			if (chosenHorse == horses[3].name) {
				if horses[3].experience == 0{
					percent := bet / 2.5
					balance += bet + percent
					fmt.Printf("%s %s %.2f", "\nYou've Won: ", "$", bet + percent)
				} else if horses[3].experience > 0 && horses[3].experience < 3 {
					percent := bet / 5
					balance += bet + percent
					fmt.Printf("%s %s %.2f", "\nYou've Won: ", "$", bet + percent)
				} else if horses[3].experience >= 3 {
					percent := bet / 10
					balance += bet + percent
					fmt.Printf("%s %s %.2f", "\nYou've Won: ", "$", bet + percent)
				}
			}
		case horses[4].name: 
			if (chosenHorse == horses[4].name) {
				if horses[4].experience == 0{
					percent := bet / 2.5
					balance += bet * 2 + percent
					fmt.Printf("%s %s %.2f", "\nYou've Won: ", "$", bet + bet + percent)
				} else if horses[4].experience > 0 && horses[4].experience < 3 {
					percent := bet / 5
					balance += bet * 2 + percent
					fmt.Printf("%s %s %.2f", "\nYou've Won: ", "$", bet + bet + percent)
				} else if horses[4].experience >= 3 {
					percent := bet / 10
					balance += bet * 2 + percent
					fmt.Printf("%s %s %.2f", "\nYou've Won: ", "$", bet + bet + percent)
				}
		}
	}	
	if (balance <= 0) {
		fmt.Printf("%s %.2f %s","\nYou have a balance of: ", balance, "\nYou can't place any more bets.")
	}
	return balance
}

func main() {
	balance := 100.00	

	for balance > 0 {
		var input string
		fmt.Println("")
		fmt.Printf("Your balance is: $%.2f ", balance) 
		fmt.Println("Enter your bet (q to quit): \n") 
		
		fmt.Scanln(&input)

		if input == "q"{
			fmt.Printf("%s %.2f","You've left with: $", balance)
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

