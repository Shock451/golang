package main

import "fmt"

func greatest(input ...int)(int){
	x:=input[0]
	for _, value:= range input{
		if value > x{
			x = value
		}
	}
	return x
}

func main(){
	fmt.Println(greatest(1,3,56,59,221,24))
}
