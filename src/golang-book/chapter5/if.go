package main
import "fmt"
/*
If

Let's modify the program we just wrote so that instead of just printing the numbers 1-10 on each line it also specifies whether or not the number is even or odd. Like this:

1 odd
2 even
3 odd
4 even
5 odd
6 even
7 odd
8 even
9 odd
10 even

First we need a way of determining whether or not a number is even or odd. An easy way to tell is to divide the number by 2. If you have nothing left over then the number is even, otherwise it's odd. So how do we find the remainder after division in Go? We use the % operator. 1 % 2 equals 1, 2 % 2 equals 0, 3 % 2 equals 1 and so on.

Next we need a way of choosing to do different things based on a condition. For that we use the if statement:

if i % 2 == 0 {
  // even
} else {
  // odd
}

An if statement is similar to a for statement in that it has a condition followed by a block. If statements also have an optional else part. If the condition evaluates to true then the block after the condition is run, otherwise either the block is skipped or if the else block is present that block is run.

If statements can also have else if parts:

if i % 2 == 0 {
  // divisible by 2
} else if i % 3 == 0 {
  // divisible by 3
} else if i % 4 == 0 {
  // divisible by 4
}

The conditions are checked top down and the first one to result in true will have its associated block executed. None of the other blocks will execute, even if their conditions also pass. (So for example the number 8 is divisible by both 4 and 2, but the // divisible by 4 block will never execute because the // divisible by 2 block is done first)

Putting it all together we have:
*/
func main() {
  for i := 1; i <= 10; i++ {
    if i % 2 == 0 {
      fmt.Println(i, "even")
    } else {
      fmt.Println(i, "odd")
    }
  }
}
/*
Let's walk through this program:

    Create a variable i of type int and give it the value 1

    Is i less than or equal to 10? Yes: jump to the block

    Is the remainder of i ÷ 2 equal to 0? No: jump to the else block

    Print i followed by odd

    Increment i (the statement after the condition)

    Is i less than or equal to 10? Yes: jump to the block

    Is the remainder of i ÷ 2 equal to 0? Yes: jump to the if block

    Print i followed by even

    …

The remainder operator, while rarely seen outside of elementary school, turns out to be really useful when programming. You'll see it turn up everywhere from zebra striping tables to partitioning data sets.
*/
