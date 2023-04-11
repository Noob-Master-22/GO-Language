package main

import "fmt"

func main() {
	fmt.Println("If else in go")

	age := 18

	if age > 18 {
		fmt.Println("You are eligible to vote")

	} else if age < 18 {
		fmt.Println("Choti Bachi Ho ")

	} else {
		fmt.Println("Fresh outta the oven")
	}

	//declaring,initializing,checking for a condition in the same line

	if num := 18; num > 10 {
		fmt.Println("Num is greater than 10")
	} else {
		fmt.Println("Num is less than 10")
	}

}
