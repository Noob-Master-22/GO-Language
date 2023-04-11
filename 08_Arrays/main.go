package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in go")

	var albums [5]string

	albums[0] = "Taylor Swift"
	albums[1] = "Fearless"
	albums[2] = "Speak Now"
	albums[3] = "Red"
	albums[4] = "1989"

	var songs = [2]string{"Ciwyw", "Dwoht"}

	fmt.Println("Songs List is", songs)
	fmt.Println("Songs Length is", len(songs))

	fmt.Println("Albums List is", albums)
	fmt.Println("Album Length is", len(albums))

}
