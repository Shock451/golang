/*

************************************                                                *****************************************
                                       Program prints out factors of 3 in 1-100                                              
                                                Abdullahi Oladosu                                                            
                                                        Go                                                                   
*/

package main
import "fmt"

func main(){
        for i:=1; i<=100; i++{
                if (i%3 == 0){
                        fmt.Println(i)
                }
        }


}
