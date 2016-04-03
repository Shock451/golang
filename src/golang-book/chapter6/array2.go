/*

The Go compiler won't allow you to create variables that you never use. Since we don't use i inside of our loop we need to change it to this:

var total float64 = 0
for _, value := range x {
  total += value
}
fmt.Println(total / float64(len(x)))

A single _ (underscore) is used to tell the compiler that we don't need this. (In this case we don't need the iterator variable)

Go also provides a shorter syntax for creating arrays:

x := [5]float64{ 98, 93, 77, 82, 83 }

We no longer need to specify the type because Go can figure it out. Sometimes arrays like this can get too long to fit on one line, so Go allows you to break it up like this:

x := [5]float64{
  98,
  93,
  77,
  82,
  83,
}

Notice the extra trailing , after 83. This is required by Go and it allows us to easily remove an element from the array by commenting out the line:

x := [4]float64{
  98,
  93,
  77,
  82,
  // 83,
}

*/
