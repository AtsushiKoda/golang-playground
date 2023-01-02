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

	fmt.Println("struct-pointers.go:")
	structPointers()

	fmt.Println("slices-pointers.go:")
	slicesPointers()

	fmt.Println("slice-len-cap.go:")
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)
	s = s[:0]
	printSlice(s)
	s = s[:4]
	printSlice(s)
	s = s[2:]
	printSlice(s)

	fmt.Println("slices-of-slice.go:")
	ticTacToe()

	fmt.Println("map-literals.go:")
	fmt.Println(m)

	fmt.Println("mutating-maps.go:")
	mutateMaps()

	fmt.Println("function-values.go:")
	funcValues()

	fmt.Println("function-closures.go:")
	funcClosures()
}
