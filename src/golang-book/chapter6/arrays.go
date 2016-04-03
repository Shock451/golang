/*

Arrays, Slices and Maps

In chapter 3 we learned about Go's basic types. In this chapter we will look at three more built-in types: arrays, slices and maps.

                                         Arrays

An array is a numbered sequence of elements of a single type with a fixed length. In Go they look like this:

var x [5]int

x is an example of an array which is composed of 5 ints. Try running the following program:

*/

package main

import "fmt"

func main() {
  var x [5]int
  x[4] = 100
  fmt.Println(x)
}
/*
You should see this:

[0 0 0 0 100]

x[4] = 100 should be read “set the 5th element of the array x to 100”. It might seem strange that x[4] represents the 5th element instead of the 4th but like strings, arrays are indexed starting from 0. Arrays are accessed in a similar way. We could change fmt.Println(x) to fmt.Println(x[4]) and we would get 100.

Here's an example program that uses arrays:

func main() {
  var x [5]float64
  x[0] = 98
  x[1] = 93
  x[2] = 77
  x[3] = 82
  x[4] = 83

  var total float64 = 0
  for i := 0; i < 5; i++ {
    total += x[i]
  }
  fmt.Println(total / 5)
}

This program computes the average of a series of test scores. If you run it you should see 86.6. Let's walk through the program:

    First we create an array of length 5 to hold our test scores, then we fill up each element with a grade

    Next we setup a for loop to compute the total score

    Finally we divide the total score by the number of elements to find the average

This program works, but Go provides some features we can use to improve it. First these 2 parts: i < 5 and total / 5 should throw up a red flag for us. Say we changed the number of grades from 5 to 6. We would also need to change both of these parts. It would be better to use the length of the array instead:

var total float64 = 0
for i := 0; i < len(x); i++ {
  total += x[i]
}
fmt.Println(total / len(x))
*/
