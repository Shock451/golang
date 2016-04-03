/*
Returning Multiple Values

Go is also capable of returning multiple values from a function:

*/

package main

import "fmt"

func f() (int, int) {
  return 5, 6
}

func main() {
  x, y := f()
  fmt.Println(x,y)
}

/*
Three changes are necessary: change the return type to contain multiple types separated by ,, change the expression after the return so that it contains multiple expressions separated by , and finally change the assignment statement so that multiple values are on the left side of the := or =.

Multiple values are often used to return an error value along with the result (x, err := f()), or a boolean to indicate success (x, ok := f()).

*/
