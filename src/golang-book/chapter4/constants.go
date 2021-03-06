/*
Constants

Go also has support for constants. Constants are basically variables whose values cannot be changed later. They are created in the same way you create variables but instead of using the var keyword we use the const keyword:
*/
package main

import "fmt"

func main() {
  const x string = "Hello World"
  fmt.Println(x)
}
/*
This:

const x string = "Hello World"
x = "Some other string"

Results in a compile-time error:

.\main.go:7: cannot assign to x

Constants are a good way to reuse common values in a program without writing them out each time. For example Pi in the math package is defined as a constant.

*/
