package main 

import "fmt"

func main(){

	var x int = 10 

	if x > 5 && x < 15 {
		fmt.Println("x is between 5 and 15")
	}

	if x > 5 || x < 15 {
		fmt.Println("x is either greater than 5 or less than 15")
	}

	if !(x < 5 && x < 9){
		fmt.Println("x is not less than 5 and less than 10")
	}
}