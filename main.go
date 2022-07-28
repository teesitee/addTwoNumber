package main

import "fmt"

func main() {
	x := -393
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
	lenght := len(number)
	if number[0] == number[lenght-1] {
		fmt.Println("Palindrome : true")
		return
	} else {
		fmt.Println("Palindrome : false")
	}
}
