package main 
import "fmt"

func main(){
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for j := 0; j <=100; j += 10 {
		fmt.Println(j)
	}

	for k := 0; k < 5; k++ {
		if k == 3 {
			continue
		}
		fmt.Println(k)
	}

	for l := 0; l < 5; l++ {
		if l == 2 {
			break
		}
		fmt.Println(l)
	}
	
	adj := [2]string{"big","tasty"}
	fruits := [3]string{"apple","orange","banana"}
	for m := 0; m < len(adj); m++ {
		for n := 0; n < len(fruits); n++ {
			fmt.Println(adj[m],fruits[n])
		}
	}

	for idx, val := range fruits {
		fmt.Println(idx,val)
	}

	for _, val := range fruits {
		fmt.Println(val)
	}

	for idx, _ := range fruits {
		fmt.Println(idx)
	}
}