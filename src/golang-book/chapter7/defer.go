/*

Defer, Panic & Recover

Go has a special statement called defer which schedules a function call to be run after the function completes. Consider the following example:

*/

package main

import "fmt"

func first() {
  fmt.Println("1st")
}
func second() {
  fmt.Println("2nd")
}
func main() {
  defer second()
  first()
}

/*
This program prints 1st followed by 2nd. Basically defer moves the call to second to the end of the function:

func main() {
  first()
  second()
}

defer is often used when resources need to be freed in some way. For example when we open a file we need to make sure to close it later. With defer:

f, _ := os.Open(filename)
defer f.Close()

This has 3 advantages: (1) it keeps our Close call near our Open call so it's easier to understand, (2) if our function had multiple return statements (perhaps one in an if and one in an else) Close will happen before both of them and (3) deferred functions are run even if a run-time panic occurs.

*/
