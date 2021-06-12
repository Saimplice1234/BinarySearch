package main

import "fmt"

func main(){

	V := 7
	T := []int{1,2,3,4,5,6,7,9}
	A := 1
	B := len(T)
	F := false

	for F == false{
		
	    M := (A+B)/2
		if T[M] == V {
			F = true
		} else if T[M]<V {
			A = M + 1
		} else if T[M]>V {
			B = M - 1
		}
      
		if F == true {
			fmt.Printf("Found at index %d", M)
		}

	}

}
