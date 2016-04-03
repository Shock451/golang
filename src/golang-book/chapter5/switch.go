package main
import "fmt"

func main(){
/*

Switch

Suppose we wanted to write a program that printed the English names for numbers. Using what we've learned so far we might start by doing this:

if i == 0 {
  fmt.Println("Zero")
} else if i == 1 {
  fmt.Println("One")
} else if i == 2 {
  fmt.Println("Two")
} else if i == 3 {
  fmt.Println("Three")
} else if i == 4 {
  fmt.Println("Four")
} else if i == 5 {
  fmt.Println("Five")
}

Since writing a program in this way would be pretty tedious Go provides another statement to make this easier: the switch statement. We can rewrite our program to look like this:
*/
        i:=4
        switch i {
        case 0: fmt.Println("Zero")
        case 1: fmt.Println("One")
        case 2: fmt.Println("Two")
        case 3: fmt.Println("Three")
        case 4: fmt.Println("Four")
        case 5: fmt.Println("Five")
        default: fmt.Println("Unknown Number")
}
/*
A switch statement starts with the keyword switch followed by an expression (in this case i) and then a series of cases. The value of the expression is compared to the expression following each case keyword. If they are equivalent then the statement(s) following the : is executed.

Like an if statement each case is checked top down and the first one to succeed is chosen. A switch also supports a default case which will happen if none of the cases matches the value. (Kind of like the else in an if statement)

These are the main control flow statements. Additional statements will be explored in later chapters.
*/
}
