package main

import "fmt"


func main(){
	age:=18
	sex:="male"

	if age==20 && sex=="male"{
		fmt.Println("You are applicable as a male")
		}else if age<=17 && sex=="female"{
		fmt.Println("You are applicable as a female")
		
		}else if age==18 || sex=="female"{
			fmt.Println("you are good to go");
		}else{
		fmt.Println("You are not applicable ")

	}

	


}

//  >,>=,<,<=,==
// and => &&
//or => ||
//not => !