package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("named-results.go:")
	fmt.Println(split(17))

	fmt.Println("basic-types.go:")
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	fmt.Println("if-with-a-short-statement.go:")
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	default:
		fmt.Println("Too far away.")
	}

	fmt.Println("defer-multi.go:")
	stackDefer()
}
