package main 

import "fmt"

func main(){
	var x int = 5 
	var y int = 3 

	if x > y {
		fmt.Println("x is greater than y")
	}

	if x < y {
		fmt.Println("x is less than y")
	}
	
	if x == y {
		fmt.Println("x is equal to y")
	}

	if x != y {
		fmt.Println("x is not equal to y")
	}

	if x >= y {
		fmt.Println("x is greater than or equal to y")
	}

	if x <= y {
		fmt.Println("x is less than or equal to y")
	}
}