//Other programming languages have a lot of different types of loops (while, do, until, foreach, …) but Go only has one that can be used in a variety of different ways. The previous program could also have been written like this:

func main() {
  for i := 1; i <= 10; i++ {
    fmt.Println(i)
  }
}
/*
Now the conditional expression also contains two other statements with semicolons between them. First we have the variable initialization, then we have the condition to check each time and finally we “increment” the variable. (adding 1 to a variable is so common that we have a special operator: ++. Similarly subtracting 1 can be done with --)

We will see additional ways of using the for loop in later chapters.
*/
