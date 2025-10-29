package main

import "fmt"


func main(){
	a:=4

	switch a{
	case 1:
		fmt.Println(1)
	case 2, 4:	
		fmt.Println("2 and 4")
	default:
		fmt.Println(3)	
	}
}

