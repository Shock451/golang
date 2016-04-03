/* Program that converts from fahrenheit to celsius

Takes in input as fahrenheit, converts and outputs it in celsius

*/

package main

import "fmt"

func f_to_c(f float64){
        c:= (f - 32.0) * (5.0/9.0)
        fmt.Print("The result is ", c, "C\n")
}

func main(){

        var input float64
        print("What's the temperature in fahrenheit: ")
        fmt.Scanf("%f", &input)
        f_to_c(input)
}
