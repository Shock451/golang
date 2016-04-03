/*

Functions

A function is an independent section of code that maps zero or more input parameters to zero or more output parameters. Functions (also known as procedures or subroutines) are often represented as a black box: (the black box represents the function)

Until now the programs we have written in Go have used only one function:

func main() {}

We will now begin writing programs that use more than one function.
Your Second Function

Remember this program from chapter 6:

func main() {
  xs := []float64{98,93,77,82,83}

  total := 0.0
  for _, v := range xs {
    total += v
  }
  fmt.Println(total / float64(len(xs)))
}

This program computes the average of a series of numbers. Finding the average like this is a very general problem, so it's an ideal candidate for definition as a function.

The average function will need to take in a slice of float64s and return one float64. Insert this before the main function:

func average(xs []float64) float64 {
  panic("Not Implemented")
}

Functions start with the keyword func, followed by the function's name. The parameters (inputs) of the function are defined like this: name type, name type, …. Our function has one parameter (the list of scores) that we named xs. After the parameters we put the return type. Collectively the parameters and the return type are known as the function's signature.

Finally we have the function body which is a series of statements between curly braces. In this body we invoke a built-in function called panic which causes a run time error. (We'll see more about panic later in this chapter) Writing functions can be difficult so it's a good idea to break the process into manageable chunks, rather than trying to implement the entire thing in one large step.

Now let's take the code from our main function and move it into our average function:
*/

func average(xs []float64) float64 {
  total := 0.0
  for _, v := range xs {
    total += v
  }
  return total / float64(len(xs))
}
/*
Notice that we changed the fmt.Println to be a return instead. The return statement causes the function to immediately stop and return the value after it to the function that called this one. Modify main to look like this:

*/

func main() {
  xs := []float64{98,93,77,82,83}
  fmt.Println(average(xs))
}

