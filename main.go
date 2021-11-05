package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Welcome to GO Calculator!")
	cmd := readLine("Enter command: [a]dd, [s]ubtract, [m]ultiply, [d]ivide:")
	fmt.Println(cmd)
	/*if cmd == "a" {
		num1, num2 := getUserNumbers()
		result := num1 + num2 //cannot just print an int so you need to convert it back to a string!
		sResult := fmt.Sprintf("%d", result)
		fmt.Println(sResult)

	} else if cmd == "s" {
		num1, num2 := getUserNumbers()
		result := num1 - num2 //cannot just print an int so you need to convert it back to a string!
		sResult := fmt.Sprintf("%d", result)
		fmt.Println(sResult)

	} else if cmd == "m" {
		num1, num2 := getUserNumbers()
		result := num1 * num2 //cannot just print an int so you need to convert it back to a string!
		sResult := fmt.Sprintf("%d", result)
		fmt.Println(sResult)

	} else if cmd == "d" {
		num1, num2 := getUserNumbers()
		result := float32(num1) / float32(num2)
		sResult := fmt.Sprintf("%f", result)
		fmt.Println(sResult)

	} else {
		fmt.Println("Invalid Input!")
	} */

	switch(cmd){
	case "a":
		num1, num2 := getUserNumbers()
		result := num1 + num2 //cannot just print an int so you need to convert it back to a string!
		sResult := fmt.Sprintf("%d", result)
		fmt.Println(sResult)
		
	case "b":
		num1, num2 := getUserNumbers()
		result := num1 - num2 //cannot just print an int so you need to convert it back to a string!
		sResult := fmt.Sprintf("%d", result)
		fmt.Println(sResult)
		
	case "c":
		num1, num2 := getUserNumbers()
		result := num1 * num2 //cannot just print an int so you need to convert it back to a string!
		sResult := fmt.Sprintf("%d", result)
		fmt.Println(sResult)
		
	case "d":
		num1, num2 := getUserNumbers()
		result := float32(num1) / float32(num2)
		sResult := fmt.Sprintf("%f", result)
		fmt.Println(sResult)
		
	 default:
		fmt.Println("Invalid input!")

}

func readLine(message string) string {
	fmt.Print(message)
	var input string
	fmt.Scanln(&input)
	return input
}

func getUserNumbers() (int, int) {
	num1String := readLine("First Number:")
	num1, err := strconv.Atoi(num1String) //turns the string into an int
	if err != nil {
		fmt.Println(err)
	}

	num2String := readLine("Second Number:")
	num2, err := strconv.Atoi(num2String) //turns the string into an int
	if err != nil {
		fmt.Println(err)
	}

	return num1, num2
}

