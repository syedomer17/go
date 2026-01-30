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

	// go slicing and capacity of array

	numbers := [6]int{10,11,12,13,14,15}

	newNumber := numbers[2:4]

	fmt.Printf("%v\n",numbers)
	fmt.Println(newNumber)
	fmt.Printf("%d\n",len(newNumber))
	fmt.Printf("%d\n",cap(newNumber))

	mySlice := make([]int,5,10)
	fmt.Println(mySlice)
	mySlice[0] = 100
	mySlice[1] = 200
	mySlice[2] = 300
	mySlice[3] = 400
	mySlice[4] = 500
	// mySlice[5] = 600  // this will cause runtime error: index out of range
	fmt.Println(mySlice)
	fmt.Printf("Length: %d\n",len(mySlice))
	fmt.Printf("Capacity: %d\n",cap(mySlice))


	// append function to add elements to slice
	mySlice2 := []int{1,2,3,4,5,6}
	fmt.Println(mySlice2)
	fmt.Printf("Length: %d\n",len(mySlice2))
	fmt.Printf("Capacity: %d\n",cap(mySlice2))

	mySlice3 := append(mySlice2, 20,21)
	fmt.Println(mySlice3)
	fmt.Printf("Length: %d\n",len(mySlice3))
	fmt.Printf("Capacity: %d\n",cap(mySlice3))
}