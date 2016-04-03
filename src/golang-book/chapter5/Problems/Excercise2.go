/*
************************************                                                *****************************************
                                              Program prints out 1-100                                              
                                       ....Prints Fizz for multiples of 3....................................................
                                       ....Prints Buzz for multiples of 5....................................................
                                       Prints FizzBuzz for multiples of 3 and 5..............................................
                                                Abdullahi Oladosu                                                            
                                                        Go                                                                   
*/
package main
import "fmt"

func main(){
        for i:=0; i<=100; i++{
                if (i%3 == 0 && i%5 == 0){
                        fmt.Println("FizzBuzz")
                }else if (i%3 == 0){
                        fmt.Println("Fizz")
                }else if(i%5 == 0) { 
                        fmt.Println("Buzz")
                }else {
                        fmt.Println(i)
                }
        }
}
