/*

Here's an example program which takes in a number entered by the user and doubles it:

We use another function from the fmt package to read the user input (Scanf). &input will be explained in a later chapter, for now all we need to know is that Scanf fills input with the number we enter.
*/
package main

import "fmt"

func main() {
  fmt.Print("Enter a number: ")
  var input float64
  fmt.Scanf("%f", &input)

  output := input * 2

  fmt.Println(output)
}


