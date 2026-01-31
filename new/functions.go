package main 
import "fmt"

func myMessage() {
	fmt.Println("I just got executed!")
}

func familyName(fname string) {
	fmt.Println("Hello", fname)
}

func ageWithName(fname string, age int){
	fmt.Println("Hello", fname, "You are", age, "Years old")
}

func add(a int, b int) int {
	return a + b
}

func newadd(x int, y int) (result int){
	result = x + y 
	return
}

func newaddfunc(x int, y, int) (result int){
	result = x + y 
	return result
}

func myFunc(x int y string) (result int, text1 string){
	result = x + x 
	text1 = y + "World!"
	return
}

func main(){
	myMessage()
	familyName("Omer")
	familyName("Amjad")

	ageWithName("Omer", 19)

	fmt.Println(add(1,2))
	fmt.Println(newadd(3,4))
	fmt.Println(newaddfunc(5,10))

	_, b := myFunc(5,"Hello")
	fmt.Println(b)
}