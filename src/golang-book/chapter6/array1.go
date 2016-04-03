/*

Go ahead and make these changes and run the program. You should get an error:

$ go run tmp.go
# command-line-arguments
.\tmp.go:19: invalid operation: total / 5 (mismatched types float64 and int)

The issue here is that len(x) and total have different types. total is a float64 while len(x) is an int. So we need to convert len(x) into a float64:

fmt.Println(total / float64(len(x)))

This is an example of a type conversion. In general to convert between types you use the type name like a function.

Another change to the program we can make is to use a special form of the for loop:

var total float64 = 0
for i, value := range x {
  total += value
}
fmt.Println(total / float64(len(x)))

In this for loop i represents the current position in the array and value is the same as x[i]. We use the keyword range followed by the name of the variable we want to loop over.

Running this program will result in another error:

$ go run tmp.go
# command-line-arguments
.\tmp.go:16: i declared and not used

*/
