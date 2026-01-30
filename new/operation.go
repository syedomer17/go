package main
import ("fmt")

func main(){
	var x int = 10
	var y int = 20 

	sum := x + y 
	fmt.Printf("Sum of %d and %d is %d\n", x, y, sum)

	diff := y - x
	fmt.Printf("Difference of %d and %d is %d\n",y,x,diff)

	prod := x * y 
	fmt.Printf("Product of %d and %d is %d\n",x,y,prod)

	quot := y / x 
	fmt.Printf("Quotient of %d and %d is %d\n",y,x,quot)

	mod := y % x 
	fmt.Printf("Modulus of %d and %d is %d\n",y,x,mod)

	x++
	fmt.Printf("Value of x after increment is %d\n",x)

	y--
	fmt.Printf("Value of y after decrement is %d\n",y)
}