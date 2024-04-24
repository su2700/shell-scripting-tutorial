package main

import "fmt"

func main() {
	var X int
	var Y int
	fmt.Println(
        "add of X and Y")
    fmt.Print("enter X: ")
    _, err := fmt.Scan(&X)
	fmt.Println(X)
	if err != nil {
		fmt.Println(err)
		return
	}


}