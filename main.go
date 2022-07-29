package main

import "fmt"

func main() {
	x := 11211
	if x <= 0 {
		fmt.Println("Palindrome : false")
		return
	}
	forCount := x
	count := 0

	for forCount > 0 {
		forCount /= 10
		count++
	}
	reCount := count
	var number []int
	var s int
	var s2 int
	devine := 1
	counter := count - 1
	for counter > 0 {
		counter--
		devine *= 10
	}
	num := x
	s = num
	for count > 0 {
		s /= devine
		count--
		number = append(number, s)
		s2 = num % devine
		s = s2
		devine /= 10
	}
	fmt.Println(number)

	var reNumber []int
	for i := range number {
		reNumber = append(reNumber, number[reCount-1-i])
	}
	fmt.Println("renumber :", reNumber)
	for i := range number {
		if number[i] == reNumber[i] {
			continue
		} else {
			fmt.Println("Palindrome : false")
			return
		}
	}
	fmt.Println("Palindrome : true")

}
