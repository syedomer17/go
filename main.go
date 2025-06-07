package main

import "fmt"

func main(){
	fmt.Println("Hello,Go is working!")

	var name string = "Omer"
    var age int = 18
    height := 5.9  // short declaration
    isCoder := true

    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Height:", height)
    fmt.Println("Is a Coder:", isCoder)

	var a = 2
	var b = 3
	var sum = a + b
	fmt.Println("sum:",sum)

	if age>= 18 {
		fmt.Println("✅ You are an adult.")
	}else {
		fmt.Println("❌ You are a minor.")
	}

	fmt.Println("🔢 Numbers from 1 to 5:")
	for i :=1; i <= 5; i++ {
		fmt.Println(i)
	}

	for i :=1; i <=20; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		}else if i%3 == 0 {
			 fmt.Println("Fizz")
		}else if i%5 == 0 {
			fmt.Println("Buzz")
		}else {
			fmt.Println(1)
		}
	}
}
