package main
import "fmt"

func main(){

	x := 40
	y := 20 

	if x > y {
		fmt.Println("x is greater than y")
	}else if x < y {
		fmt.Println("x is less than y")
	}else {
		fmt.Println("x is equal to y")
	}

	marks := 39

	if marks >= 40 { // pass condition
    if marks >= 90 {
        fmt.Println("Grade: A+")
    } else if marks >= 80 {
        fmt.Println("Grade: A")
    } else if marks >= 70 {
        fmt.Println("Grade: B")
    } else if marks >= 60 {
        fmt.Println("Grade: C")
    } else if marks >= 50 {
        fmt.Println("Grade: D")
    } else {
        fmt.Println("Grade: E")
    }
} else {
    fmt.Println("Fail")
}

}