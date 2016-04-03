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
        var x,y bool
        for i:=0; i<=100; i++{ 
                x = i%3 == 0 
                y = i%5 == 0
                switch(true){
                        case x && y: fmt.Println("FizzBuzz") 
                        case x && !y: fmt.Println("Fizz")
                        case !x && y: fmt.Println("Buzz")
                        case !x && !y: fmt.Println(i)
                }
        }
}
