package main

import (
	"fmt"
	"time"
)

func greet(name string)  {
		fmt.Println("Hello:",name)
	}

func add(a int, b int) int{
	return a + b
}

func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("❌ Cannot divide by zero")
    }
    return a / b, nil
}

type User struct {
	Name string
	Age int
	Location string
}

func (u User) Greet() {
	fmt.Println(u.Name)
}

type Speaker interface {
    Speak()
}

func (u User) Speak() {
    fmt.Println("Hi, my name is", u.Name)
}

func printMessage(msg string){
	for i :=1; i < 3; i++ {
		fmt.Println(msg)
		time.Sleep(time.Second)
	}
}

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
	greet("omer")
	greet("go devloper")

	result := add(10,20)
	fmt.Println("sum:",result)

	res, err :=  divide(10,0)
	if err != nil {
		fmt.Println("Error:",err)
	}else {
		fmt.Println("Result:",res)
	}

	var nums [3]int = [3]int{10,20,30}
	fmt.Println("Array:" ,nums)
	fmt.Println("First Element of Array:",nums[0])

	numbers := []int{1,2,3,4,5}
	numbers = append(numbers, 6)
	
	fmt.Println("slice:",numbers)
	fmt.Println("length:",len(numbers))

	user := map[string]string{
		"name":"omer",
		"lang":"go",
		"city":"hyderabad",
	}

	fmt.Println("User:",user)
	fmt.Println("name:",user["name"])

	omer := User{
		Name: "omer",
		Age: 18,
		Location: "Hyderabad",
	}
	fmt.Println(omer)
	fmt.Println(omer.Age)
	fmt.Println(omer.Name)
	fmt.Println(omer.Location)

	u := User{Name: "omer",Age: 18,Location: "HYD"}
	u.Greet()

	var s Speaker
	s= User{Name: "omer ali"}
	s.Speak()

	// go printMessage("👋 From Goroutine") // runs in background
    printMessage("📢 From Main")
}
