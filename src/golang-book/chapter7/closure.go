/*

Closure

It is possible to create functions inside of functions:

*/

func main() {
  add := func(x, y int) int {
    return x + y
  }
  fmt.Println(add(1,1))
}

/*
add is a local variable that has the type func(int, int) int (a function that takes two ints and returns an int). When you create a local function like this it also has access to other local variables (remember scope from chapter 4):

func main() {
  x := 0
  increment := func() int {
    x++
    return x
  }
  fmt.Println(increment())
  fmt.Println(increment())
}

increment adds 1 to the variable x which is defined in the main function's scope. This x variable can be accessed and modified by the increment function. This is why the first time we call increment we see 1 displayed, but the second time we call it we see 2 displayed.

A function like this together with the non-local variables it references is known as a closure. In this case increment and the variable x form the closure.

One way to use closure is by writing a function which returns another function which – when called – can generate a sequence of numbers. For example here's how we might generate all the even numbers:

func makeEvenGenerator() func() uint {
  i := uint(0)
  return func() (ret uint) {
    ret = i
    i += 2
    return
  }
}
func main() {
  nextEven := makeEvenGenerator()
  fmt.Println(nextEven()) // 0
  fmt.Println(nextEven()) // 2
  fmt.Println(nextEven()) // 4
}

makeEvenGenerator returns a function which generates even numbers. Each time it's called it adds 2 to the local i variable which – unlike normal local variables – persists between calls.

*/
