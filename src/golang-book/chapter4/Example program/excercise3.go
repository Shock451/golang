/* 

***************                             *******************
                converts from feet to metre
                        oladosu                                
                         golang
                         
*/

package main

import "fmt"

func f_to_m(f float64){
        m:= 0.3048 * f
        fmt.Print("The result is ", m, "m\n")
}

func main(){

        var input float64
        print("What's the distance in feet: ")
        fmt.Scanf("%f", &input)
        f_to_m(input)
}
