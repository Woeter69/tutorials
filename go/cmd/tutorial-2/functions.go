package main

import (
	"fmt"
	"errors"
)

func main() {
	var printValue string = "We cool"
	printMe(printValue) // Have to call the function in main for getting it to terminal

	const numerator int = 4
	const denominator int = 2
	var result, remainder, err = intDivision(numerator, denominator) // Can recieve multiple variables from functions
	if err != nil { // Checking if the returned error is nil
		fmt.Printf(err.Error()) // Prints the actual error 
	}else if remainder == 0 {
		fmt.Printf("Result is %v",result)
	}else {
		fmt.Printf("Result is %v and Remainder is %v",result,remainder) 
	}

	switch{ // switch also work in Go you can replicate the above if-else statements in switch too 
		case err!=nil:
			fmt.Printf(err.Error())
		
		case remainder==0:
			fmt.Printf("Result is %v",result)

		default:
			fmt.Printf("Result is %v and Remainder is %v",result,remainder) 
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int,int,error) {  // Can return and recieve multiple values
	var err error // error is a new datatype in Go which stores and used to return Errors
	if denominator == 0 {
		err = errors.New("Can't divide by 0")
		return 0, 0, err // Passing the error is a good practice in Go 
	}

	var result int = numerator/denominator
	var remainder int = numerator%denominator
	return result, remainder, err // Most of the modern imports like modules and functions return errors as a parameter in Go
}
