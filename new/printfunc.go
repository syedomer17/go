//there are three functions to output text

/*
Print() function prints its arguments with their default format
Println() function prints its arguments with their default format and adds a newline at the end
Printf() function prints its arguments according to a specified format
*/

package main 
import ("fmt")

func main(){
	var i,j string = "Hello","World"

	fmt.Print(i,j)
	fmt.Println(i,j)
	fmt.Printf("i has value: %v and type: %T\n",i,i)
	fmt.Printf("j has value: %v and type: %T\n",j,j)
}