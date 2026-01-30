// there are two ways to declare and create arrays in Go 

package main 

import ("fmt")

func main(){
	var arr1 = [3]int{1,2,3}
	arr2 := [5]int{4,5,6,7,8}

	fmt.Println(arr1)
	fmt.Println(arr2)

	var newArr = [...]int{1,2,3,10}
	arr3 := [...]int{4,5,6,7,8,9}

	fmt.Println(newArr)
	fmt.Println(arr3)

	price := [3]int{100,200,300}
	fmt.Println(price[0])
	fmt.Println(price[1])
	fmt.Println(price[2])

	price[2] = 500 
	fmt.Println(price)

	myarr := [5]int{}
	myarr2 := [3]int{1,2}
	myarr3 := [4]string{"Go", "is", "awesome"}

	fmt.Println(myarr)
	fmt.Println(myarr2)
	fmt.Println(myarr3)
	fmt.Printf("%q\n", myarr3[3])

	myarr4 := [5]int{1:10,2:40}
	fmt.Println(myarr4)

	fmt.Println(len(myarr3))

}