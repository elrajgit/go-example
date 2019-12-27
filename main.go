package main

import "fmt"

func main() {

	sliceInts := []int{}

	for i := 0; i <= 10; i++ {
		sliceInts = append(sliceInts, i)
	}

	for _, e := range sliceInts {
		if e%2 == 0 {
			fmt.Println(e, "is even number")
		} else {
			fmt.Println(e, "is odd number")
		}
	}

}
