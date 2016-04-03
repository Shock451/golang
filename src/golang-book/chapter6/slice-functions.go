/*
Slice Functions

Go includes two built-in functions to assist with slices: append and copy. Here is an example of append:

func main() {
  slice1 := []int{1,2,3}
  slice2 := append(slice1, 4, 5)
  fmt.Println(slice1, slice2)
}

After running this program slice1 has [1,2,3] and slice2 has [1,2,3,4,5]. append creates a new slice by taking an existing slice (the first argument) and appending all the following arguments to it.

Here is an example of copy:

func main() {
  slice1 := []int{1,2,3}
  slice2 := make([]int, 2)
  copy(slice2, slice1)
  fmt.Println(slice1, slice2)
}

After running this program slice1 has [1,2,3] and slice2 has [1,2]. The contents of slice1 are copied into slice2, but since slice2 has room for only two elements only the first two elements of slice1 are copied.
*/
