package main

import "fmt"

func main() {
	var numbers []int

	for i := 1; i < 101; i++ {
		numbers = append(numbers, i)
	}

	for _, ele := range numbers {

		if ele%2 == 0 {
			fmt.Println(ele, " is an Even Number.")
		} else {
			fmt.Println(ele, " is an Odd Number.")
		}
	}
}
