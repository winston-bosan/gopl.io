package main

import "fmt"

func main() {
	fmt.Println(panic_recover(34))
}

func panic_recover(zero int) (result int) {
	defer func() {
		switch p:=recover(); p{
			default:
				fmt.Println(33)
				result = 3
		}
	}()
	panic("OMG")
}