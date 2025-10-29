package main

import "fmt"


func main(){
	// way-1
	// var x int =10
	// var name string="Rayied"
	// var isOk bool=false;
	
	//way-2
	// var a =10
	a:=40.34
	b:="hello world"
	fmt.Println((a))
	fmt.Println((b))
	c:=true
	fmt.Println((c))
	c=false
	fmt.Println((c))
	// c="asasa" #error
	// fmt.Println((c))


	const p=100

	// p=10 //error
	fmt.Println((p))



}
// variable is a container where we can store some data.
// datatypes-> numeric(number->int,float),boolean(T/F):bool,string
// float-> float32,float64
// int-> int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64
// const: unchangeable
