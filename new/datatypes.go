package main

import "fmt"

func main(){
var s1 string 

s1 = "Learn Go!"

// declare multiple variables at once 
var b, c int = 1,2
var d = true 

fmt.Println(s1,b,c,d)

//short declaration 
s2 := "Learn Go!"
e, f := 1,2
g := true 

fmt.Println(s2,e,f,g)

/*
int stores integers 123 or -123 
fload32 stores floating points number 19.12
string 
bool
*/
// var can be used outside of the function
var s string = "Hello omer!"
// := cannot be used outside of the funcation
x := 2

var y float32 = 3.14

var z bool = true

fmt.Println(x,y,z,s)

var bc, ad , dc, ac int = 1, 2, 3, 4
fmt.Println(bc, ad , dc, ac )

var one,two = 6, "Golang"
three, four := 7, "Programming"

fmt.Println(one, two, three, four)

//Multiple variable declarations can also be grouped together into a block for greater readability:

var (
	h int 
	i int =1 
	j string = "hee"
)

fmt.Println(h,i,j)
}
