/*
Recursion

Finally a function is able to call itself. Here is one way to compute the factorial of a number:

*/

package main

import "fmt"

func main(){
	fmt.Println(factorial(5))
}

func factorial(x uint) uint {
  if x == 0 {
    return 1
  }
  return x * factorial(x-1)
}
/*
factorial calls itself, which is what makes this function recursive. In order to better understand how this function works, lets walk through factorial(2):

    Is x == 0? No. (x is 2)

    Find the factorial of x – 1

        Is x == 0? No. (x is 1)

        Find the factorial of x – 1

            Is x == 0? Yes, return 1.

        return 1 * 1

    return 2 * 1

Closure and recursion are powerful programming techniques which form the basis of a paradigm known as functional programming. Most people will find functional programming more difficult to understand than an approach based on for loops, if statements, variables and simple functions.

*/
