package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch Case in Go")

	rand.Seed(time.Now().UnixNano())
	Number := rand.Intn(6) + 1
	fmt.Println("Value of dice is", Number)

	switch Number {
	case 1:
		fmt.Printf("Dice Value is 1 and you can open")
	case 2:
		fmt.Println("Move 2 spots")
	case 3:
		fmt.Println("Move 3 spots")
	case 4:
		fmt.Println("Move 4 spots")
	case 5:
		fmt.Println("Move 5 spots")
	case 6:
		fmt.Println("Move 6 spots and roll your dice again")
	default:
		fmt.Println("WTF Broo!!")

	}
}
