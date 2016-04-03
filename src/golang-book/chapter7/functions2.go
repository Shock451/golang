/*
Running this program should give you exactly the same result as the original. A few things to keep in mind:

    The names of the parameters don't have to match in the calling function. For example we could have done this:

    func main() {
      someOtherName := []float64{98,93,77,82,83}
      fmt.Println(average(someOtherName))
    }

    And our program would still work.

    Functions don't have access to anything in the calling function. This won't work:

    func f() {
      fmt.Println(x)
    }
    func main() {
      x := 5
      f()
    }

    We need to either do this:

    func f(x int) {
      fmt.Println(x)
    }
    func main() {
      x := 5
      f(x)
    }

    Or this:

    var x int = 5
    func f() {
      fmt.Println(x)
    }
      func main() {
      f()
    }

    Functions are built up in a “stack”. Suppose we had this program:

    func main() {
      fmt.Println(f1())
    }
    func f1() int {
      return f2()
    }
    func f2() int {
      return 1
    }

    We could visualize it like this:

    Each time we call a function we push it onto the call stack and each time we return from a function we pop the last function off of the stack.

    We can also name the return type:

    func f2() (r int) {
      r = 1
      return
    }

*/
