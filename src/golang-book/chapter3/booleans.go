package main

import "fmt"
/*
Booleans

A boolean value (named after George Boole) is a special 1 bit integer type used to represent true and false (or on and off). Three logical operators are used with boolean values:
&&	and
||	or
!	not
Here is an example program showing how they can be used:

*/


func main() {
  fmt.Println(true && true)
  fmt.Println(true && false)
  fmt.Println(true || true)
  fmt.Println(true || false)
  fmt.Println(!true)
}

/*
Running this program should give you:

$ go run main.go
true
false
true
true
false

We usually use truth tables to define how these operators work:
Expression	Value
true && true	true
true && false	false
false && true	false
false && false	false
Expression	Value
true || true	true
true || false	true
false || true	true
false || false	false
Expression	Value
!true	false
!false	true

These are the simplest types included with Go and form the foundation from which all later types are built. 

*/
