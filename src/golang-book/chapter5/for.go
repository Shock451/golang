/*
For

The for statement allows us to repeat a list of statements (a block) multiple times. Rewriting our previous program using a for statement looks like this:


*/

package main

import "fmt"

func main() {
  i := 1
  for i <= 10 {
    fmt.Println(i)
    i = i + 1
  }
}
/*
First we create a variable called i that we use to store the number we want to print. Then we create a for loop by using the keyword for, providing a conditional expression which is either true or false and finally supplying a block to execute. The for loop works like this:

    We evaluate (run) the expression i <= 10 (“i less than or equal to 10”). If this evaluates to true then we run the statements inside of the block. Otherwise we jump to the next line of our program after the block. (in this case there is nothing after the for loop so we exit the program)

    After we run the statements inside of the block we loop back to the beginning of the for statement and repeat step 1.

The i = i + 1 line is extremely important, because without it i <= 10 would always evaluate to true and our program would never stop. (When this happens this is referred to as an infinite loop)

As an exercise let's walk through the program like a computer would:

    Create a variable named i with the value 1

    Is i <= 10? Yes.

    Print i

    Set i to i + 1 (i now equals 2)

    Is i <= 10? Yes.

    Print i

    Set i to i + 1 (i now equals 3)

    …

    Set i to i + 1 (i now equals 11)

    Is i <= 10? No.

    Nothing left to do, so exit
*/

