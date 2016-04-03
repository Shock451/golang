package main

import "fmt"

func main(){
	fmt.Println(half(100))
}

func half(input int)(int, bool){
	switch (input/2)%2{
		case 0: {return (input/2), true}
		default: {return (input/2), false}
	}
}
