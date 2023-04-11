package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices")

	var albums = []string{}
	fmt.Printf("Type of Album is %T\n", albums)

	albums = append(albums, "Reputation", "1989", "Folklore")
	fmt.Println(albums)

	albums = append(albums[1:2])
	fmt.Println(albums)

	highScores := make([]int, 4)

	highScores[0] = 13
	highScores[1] = 1313
	highScores[2] = 26
	highScores[3] = 52

	highScores = append(highScores, 555, 666, 777)

	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	sort.Ints(highScores)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

}
