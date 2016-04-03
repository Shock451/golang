package main

import "fmt"

func main(){
/*

Since creating a new variable with a starting value is so common Go also supports a shorter statement:
*/
        x := "Hello World"
/*
Notice the : before the = and that no type was specified. The type is not necessary because the Go compiler is able to infer the type based on the literal value you assign the variable. (Since you are assigning a string literal, x is given the type string) The compiler can also do inference with the var statement:
*/
        var y = "Hello Baby"
/*
The same thing works for other types:
*/
        z := 5
        fmt.Println(x)
        fmt.Println(y)
        fmt.Println(z)
/*
Generally you should use this shorter form whenever possible.
*/
}
