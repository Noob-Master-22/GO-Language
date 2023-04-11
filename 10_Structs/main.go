package main

import "fmt"

//There is no inheritance,Super,parent concepts in go

func main() {
	fmt.Println("Structs in golang")

	satvik := User{"Satvik", "satvikreddy221b@gmail.com", true, 16}

	fmt.Println(satvik)
	fmt.Printf("My Details are: %+v\n", satvik)
	fmt.Printf("My Name is %v and email is %v\n", satvik.name, satvik.email)

}

type User struct {
	name   string
	email  string
	status bool
	Age    int
}
