package main 
import "fmt"

type Person struct {
	name string 
	age int 
	job string 
	salary int
}

type newPerson struct {
	name string 
	age int 
	job string 
	salary int 
}



func main(){
	var per1 Person
	per1.name = "Omer"
	per1.age = 19
	per1.job = "software Engineer"
	per1.salary = 7000 

	fmt.Println(per1)

	var per2 newPerson
	per2.name = "Ali"
	per2.age = 20
	per2.job = "Data Scientist"
	per2.salary = 8000

	fmt.Println(per2)
}

func printPerson(pers newPerson){
	fmt.Println("Name:", pers.name)
	fmt.Println("Age:", pers.age)
	fmt.Println("Job:", pers.job)
	fmt.Println("Salary:", pers.salary)
}